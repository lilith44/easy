package easy

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"unsafe"
)

const bitmapSize = 64

// Bitmap a bitmap is a mapping from integers to bits.
type Bitmap struct {
	sync.Mutex

	containers []uint64
}

// NewBitmap creates a new, empty bitmap.
func NewBitmap() *Bitmap {
	return new(Bitmap)
}

// NewBitmapWithContainers creates a new bitmap with given containers.
func NewBitmapWithContainers(containers []uint64) *Bitmap {
	return &Bitmap{
		containers: containers,
	}
}

// Contain checks if the bitmap contains the given value.
func (b *Bitmap) Contain(value int) bool {
	b.Lock()
	defer b.Unlock()

	return b.contains(value)
}

// Store stores the given value into the bitmap.
func (b *Bitmap) Store(value int) {
	b.Lock()
	defer b.Unlock()

	b.store(value)
}

// StoreMulti stores multi values into the bitmap.
func (b *Bitmap) StoreMulti(values ...int) {
	b.Lock()
	defer b.Unlock()

	for _, value := range values {
		b.store(value)
	}
}

// StoreNX stores the given value into the bitmap only if the bitmap does not contain the value.
func (b *Bitmap) StoreNX(value int) bool {
	b.Lock()
	defer b.Unlock()

	if !b.contains(value) {
		b.store(value)
		return true
	}
	return false
}

// Remove removes the value from the bitmap.
func (b *Bitmap) Remove(value int) {
	b.Lock()
	defer b.Unlock()

	b.remove(value)
}

// RemoveMulti removes multi values from the bitmap.
func (b *Bitmap) RemoveMulti(values ...int) {
	b.Lock()
	defer b.Unlock()

	for _, value := range values {
		b.remove(value)
	}
}

// And only stores the values stored in both bitmaps.
func (b *Bitmap) And(bitmap *Bitmap) {
	containers := make([]uint64, len(bitmap.containers))
	copy(containers, bitmap.containers)

	b.Lock()
	defer b.Unlock()

	b.and(containers)
}

// Or stores the values stored in either bitmaps.
func (b *Bitmap) Or(bitmap *Bitmap) {
	containers := make([]uint64, len(bitmap.containers))
	copy(containers, bitmap.containers)

	b.Lock()
	defer b.Unlock()

	b.or(containers)
}

// Values lists all values stored in the bitmap.
func (b *Bitmap) Values() []int {
	b.Lock()
	defer b.Unlock()

	values := make([]int, 0)
	for i := range b.containers {
		start, tmp := i*bitmapSize, 0
		for j := 0; j < bitmapSize; j++ {
			if b.containers[i] == uint64(tmp) {
				break
			}

			if b.containers[i]&(uint64(1)<<j) > 0 {
				values = append(values, start+j)
				tmp += 1 << j
			}
		}
	}
	return values
}

// FromDB implements Conversion of xorm/convert
func (b *Bitmap) FromDB(data []byte) error {
	b.Lock()
	defer b.Unlock()

	return b.unmarshal(data)
}

// ToDB implements Conversion of xorm/convert
func (b *Bitmap) ToDB() ([]byte, error) {
	b.Lock()
	defer b.Unlock()

	return b.marshal()
}

// UnmarshalJSON implements Unmarshaler of json
func (b *Bitmap) UnmarshalJSON(data []byte) error {
	b.Lock()
	defer b.Unlock()

	return b.unmarshal(data)
}

// MarshalJSON implements Marshaler of json
func (b *Bitmap) MarshalJSON() ([]byte, error) {
	b.Lock()
	defer b.Unlock()

	return b.marshal()
}

// Clear removes all values from the bitmap.
func (b *Bitmap) Clear() {
	b.Lock()
	defer b.Unlock()

	b.clear()
}

func (b *Bitmap) contains(value int) bool {
	return len(b.containers) > value/bitmapSize && (b.containers)[value/bitmapSize]&(uint64(1)<<(value%bitmapSize)) != 0
}

func (b *Bitmap) store(value int) {
	if len(b.containers) <= value/bitmapSize {
		b.containers = append(b.containers, make([]uint64, value/64-len(b.containers)+1)...)
	}

	b.containers[value/bitmapSize] |= uint64(1) << (value % bitmapSize)
}

func (b *Bitmap) remove(value int) {
	if len(b.containers) <= value/bitmapSize {
		return
	}

	b.containers[value/bitmapSize] &= ^(uint64(1) << (value % bitmapSize))
	b.reduce()
}

func (b *Bitmap) reduce() {
	if len(b.containers) == 0 {
		return
	}

	var i int
	for ; i < len(b.containers); i++ {
		if b.containers[len(b.containers)-1-i] > 0 {
			break
		}
	}

	b.containers = b.containers[:len(b.containers)-i]
}

func (b *Bitmap) and(containers []uint64) {
	min := Min(len(b.containers), len(containers))
	for i := 0; i < min; i++ {
		b.containers[i] = b.containers[i] & containers[i]
	}
	b.containers = b.containers[:min]
}

func (b *Bitmap) or(containers []uint64) {
	if len(b.containers) < len(containers) {
		b.containers = append(b.containers, make([]uint64, len(containers)-len(b.containers))...)
	} else {
		containers = append(containers, make([]uint64, len(b.containers)-len(containers))...)
	}

	for i := range b.containers {
		b.containers[i] = b.containers[i] | containers[i]
	}
	b.reduce()
}

func (b *Bitmap) unmarshal(data []byte) (err error) {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}

	if !bytes.HasPrefix(data, []byte("[")) || !bytes.HasSuffix(data, []byte("]")) {
		err = errors.New(fmt.Sprintf("can not unmarshal %s into Bitmap", data))
		return
	}

	data = bytes.Replace(data[1:len(data)-1], []byte(`"`), []byte(""), -1)
	if len(data) == 0 {
		b.containers = nil
		return nil
	}

	byteSlice := bytes.Split(data, []byte(","))
	containers := make([]uint64, len(byteSlice))
	for i := range byteSlice {
		byteSlice[i] = bytes.TrimSpace(byteSlice[i])
		if containers[i], err = strconv.ParseUint(*(*string)(unsafe.Pointer(&byteSlice[i])), 10, 64); err != nil {
			return err
		}
	}

	b.containers = containers
	return
}

func (b *Bitmap) marshal() ([]byte, error) {
	buffer := bytes.Buffer{}
	buffer.WriteByte('[')
	for i := range b.containers {
		buffer.WriteByte('"')
		buffer.WriteString(strconv.FormatUint(b.containers[i], 10))
		buffer.WriteByte('"')
		if i != len(b.containers)-1 {
			buffer.Write([]byte(", "))
		}
	}
	buffer.WriteByte(']')
	return buffer.Bytes(), nil
}

func (b *Bitmap) clear() {
	b.containers = nil
}
