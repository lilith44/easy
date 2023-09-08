package easy

import (
	"strconv"
	"testing"
)

func TestBcryptPassword(t *testing.T) {
	num := int64(10)
	for i := int64(0); i < num; i++ {
		password := strconv.FormatInt(i, 10)
		hashed, err := BcryptPassword(password)
		if err != nil {
			t.Errorf("BcryptPassword failed: %s", err)
			t.Fail()
		}

		if !BcryptCompare(hashed, password) {
			t.Errorf("BcryptCompare(%v, %v) = false, want true", hashed, password)
			t.Fail()
		}
	}
}
