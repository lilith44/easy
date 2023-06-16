package cipher

import (
	"errors"
)

const (
	AlgorithmAES = "AES"
)

const (
	ModeECB = "ECB"
)

func NewCipher(algorithm string, mode string) (Cipher, error) {
	switch algorithm {
	case AlgorithmAES:
		switch mode {
		case ModeECB:
			return newAESECB(), nil
		}
	}

	return nil, errors.New("unsupported encrypt algorithm or mode. ")
}

type Cipher interface {
	Encode([]byte, string) (string, error)
	Decode([]byte, string) (string, error)
}
