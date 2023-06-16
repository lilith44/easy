package cipher

import (
	"crypto/aes"
	"encoding/hex"
)

type aesECB struct {
}

func newAESECB() *aesECB {
	return new(aesECB)
}

func (aesECB) Encode(secret []byte, plaintext string) (string, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}

	plainBytes := padding([]byte(plaintext))
	cipherBytes := make([]byte, len(plainBytes))
	for start, end := 0, block.BlockSize(); start <= len(plaintext); {
		block.Encrypt(cipherBytes[start:end], plainBytes[start:end])
		start = start + block.BlockSize()
		end = end + block.BlockSize()
	}

	return hex.EncodeToString(cipherBytes), nil
}

func (aesECB) Decode(secret []byte, ciphertext string) (string, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return "", nil
	}

	cipherBytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	plainBytes := make([]byte, len(cipherBytes))
	for start, end := 0, block.BlockSize(); start < len(cipherBytes); {
		block.Decrypt(plainBytes[start:end], cipherBytes[start:end])
		start = start + block.BlockSize()
		end = end + block.BlockSize()
	}

	return string(unpadding(plainBytes)), nil
}
