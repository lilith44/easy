package cipher

import "crypto/aes"

func padding(plainBytes []byte) []byte {
	result := make([]byte, (len(plainBytes)/aes.BlockSize+1)*aes.BlockSize)
	copy(result, plainBytes)

	pad := byte(len(result) - len(plainBytes))
	for i := len(plainBytes); i < len(result); i++ {
		result[i] = pad
	}
	return result
}

func unpadding(cipherBytes []byte) []byte {
	var trim int
	if len(cipherBytes) > 0 {
		trim = len(cipherBytes) - int(cipherBytes[len(cipherBytes)-1])
	}
	return cipherBytes[:trim]
}
