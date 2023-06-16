package cipher

import (
	"crypto/aes"
	"encoding/hex"
	"errors"
)

type ecbCipher struct{}

type AESMode int8

const (
	// aes ECB模式
	ECBMode AESMode = 0
)

func NewCipher(mode AESMode) (Cipher, error) {
	switch mode {
	case ECBMode:
		return &ecbCipher{}, nil
	default:
		return nil, errors.New("不支持的AES模式")
	}
}

type Cipher interface {
	Encode([]byte, string) (string, error)
	Decode([]byte, string) (string, error)
}

func (ecbCipher) Encode(secret []byte, plaintext string) (string, error) {
	origData := []byte(plaintext)
	plain := padding(origData)

	blocker, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}
	encrypted := make([]byte, len(plain))
	for start, end := 0, blocker.BlockSize(); start <= len(origData); {
		blocker.Encrypt(encrypted[start:end], plain[start:end])
		start = start + blocker.BlockSize()
		end = end + blocker.BlockSize()
	}

	return hex.EncodeToString(encrypted), nil
}

func (ecbCipher) Decode(secret []byte, ciphertext string) (string, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return "", nil
	}

	encrypted, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	decrypted := make([]byte, len(encrypted))
	for start, end := 0, block.BlockSize(); start < len(encrypted); {
		block.Decrypt(decrypted[start:end], encrypted[start:end])
		start = start + block.BlockSize()
		end = end + block.BlockSize()
	}

	decrypted = unpadding(decrypted)

	return string(decrypted), nil
}

func padding(plaintext []byte) []byte {
	plaintextLength := len(plaintext)
	length := (plaintextLength + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, plaintext)
	pad := byte(len(plain) - plaintextLength)
	for i := plaintextLength; i < len(plain); i++ {
		plain[i] = pad
	}
	return plain
}

func unpadding(decrypted []byte) []byte {
	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return decrypted[:trim]
}
