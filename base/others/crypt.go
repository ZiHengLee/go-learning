package others

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
)

func Encrypt(plainText string, key string) (cipherText string, err error) {
	var block cipher.Block
	keyBytes := hashBytes(key)
	plainTextBytes := []byte(plainText)
	block, err = aes.NewCipher(keyBytes)
	if err != nil {
		return
	}

	cipherTextBytes := make([]byte, aes.BlockSize+len(plainTextBytes))
	iv := cipherTextBytes[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherTextBytes[aes.BlockSize:], plainTextBytes)
	cipherText = "crypt-" + hex.EncodeToString(cipherTextBytes)
	return
}

func Decrypt(cipherText string, key string) (plainText string, err error) {
	if len(cipherText) == 0 || len(cipherText) < 6 || cipherText[:6] != "crypt-" {
		err = errors.New("Illegal ciphertext")
		return
	}
	cipherText = string(cipherText[6:])
	var block cipher.Block
	keyBytes := hashBytes(key)
	cipherTextBytes, _ := hex.DecodeString(cipherText)
	block, err = aes.NewCipher(keyBytes)
	if err != nil {
		return
	}

	if len(cipherTextBytes) < aes.BlockSize {
		err = errors.New("Ciphertext too short")
		return
	}

	iv := cipherTextBytes[:aes.BlockSize]
	cipherTextBytes = cipherTextBytes[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)

	plainTextBytes := make([]byte, len(cipherTextBytes))
	stream.XORKeyStream(plainTextBytes, cipherTextBytes)
	plainText = string(plainTextBytes)
	return
}

func hashBytes(key string) (hash []byte) {
	h := sha1.New()
	io.WriteString(h, key)
	hashStr := hex.EncodeToString(h.Sum(nil))
	hash = []byte(hashStr)[:32]
	return
}

func SimpleIsReviewValid(review string) (setHot bool) {
	review = unicodeEmojiCode(review)
	review = strings.TrimSpace(review)
	fmt.Println(review)
	if len(review) < 10 {
		setHot = true
		return
	}
	return
}
func unicodeEmojiCode(s string) string {
	ret := ""
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) >= 3 {
			u := ""
			ret += u

		} else {
			ret += string(rs[i])
		}
	}
	return ret
}
