package easy

import (
	"fmt"
	"sync"
	"time"
)

var (
	epoch    int64 = 1288834974657
	nodeBits uint8 = 10
	stepBits uint8 = 12
)

type Snowflake struct {
	mu    sync.Mutex
	epoch time.Time
	time  int64
	node  int64
	step  int64

	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift uint8
	nodeShift uint8
}

// UniqueIdGenerator generates an unique id.
type UniqueIdGenerator func() int64

// NewSnowflake creates a new snowflake object.
func NewSnowflake(generator UniqueIdGenerator) *Snowflake {
	snowflake := new(Snowflake)
	snowflake.node = generator()
	snowflake.nodeMax = -1 ^ (-1 << nodeBits)
	if snowflake.node < 0 || snowflake.node > snowflake.nodeMax {
		panic(fmt.Sprintf("incorrect node, must be between 0 and %d", snowflake.nodeMax))
	}

	snowflake.nodeMask = snowflake.nodeMax << stepBits
	snowflake.stepMask = -1 ^ (-1 << stepBits)
	snowflake.timeShift = nodeBits + stepBits
	snowflake.nodeShift = stepBits

	currentTime := time.Now()
	snowflake.epoch = currentTime.Add(time.Unix(epoch/1e3, (epoch%1e3)*1e6).Sub(currentTime))

	return snowflake
}

// NextId generates a new, unique id.
func (s *Snowflake) NextId() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Since(s.epoch).Milliseconds()
	if now == s.time {
		s.step = (s.step + 1) & s.stepMask

		if s.step == 0 {
			for now <= s.time {
				now = time.Since(s.epoch).Milliseconds()
			}
		}
	} else {
		s.step = 0
	}
	s.time = now

	return now<<s.timeShift | s.node<<s.nodeShift | s.step
}
