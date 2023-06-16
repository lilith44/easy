package cipher

import (
	"testing"
)

func TestCipherEncode(t *testing.T) {
	secret := "thisisverysecret"
	plainText := "This is a plain text1"
	cipherText := "b51be4030614b6682c3a6407798eac64964b76e79019421130dfa9d256a4a81f"
	c, err := NewCipher(AlgorithmAES, ModeECB)
	if err != nil {
		t.Error(err)
	}
	got, err := c.Encode([]byte(secret), plainText)
	if err != nil {
		t.Error(err)
	}
	if got != cipherText {
		t.Errorf("want %v, got %v", cipherText, got)
	}
	plainText2 := "加密解密，，"
	cipherText2 := "0997d240640b2420a796c8cd82c2fe0116ca0cac483b71cf69e042cdaa6db6bd"
	got2, err := c.Encode([]byte(secret), plainText2)
	if err != nil {
		t.Error(err)
	}
	if got2 != cipherText2 {
		t.Errorf("want %v, got %v", cipherText2, got2)
	}
}

func TestCipherDecode(t *testing.T) {
	secret := "thisisverysecret"
	plainText := "This is a plain text1"
	cipherText := "b51be4030614b6682c3a6407798eac64964b76e79019421130dfa9d256a4a81f"
	c, err := NewCipher(AlgorithmAES, ModeECB)
	if err != nil {
		t.Error(err)
	}
	got, err := c.Decode([]byte(secret), cipherText)
	if err != nil {
		t.Error(err)
	}
	if got != plainText {
		t.Errorf("want %v, got %v", plainText, got)
	}

	plainText2 := "加密解密，，"
	cipherText2 := "0997d240640b2420a796c8cd82c2fe0116ca0cac483b71cf69e042cdaa6db6bd"
	got2, err := c.Decode([]byte(secret), cipherText2)
	if err != nil {
		t.Error(err)
	}
	if got2 != plainText2 {
		t.Errorf("want %v, got %v", plainText2, got2)
	}
}
