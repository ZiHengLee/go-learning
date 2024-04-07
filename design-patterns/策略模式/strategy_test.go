package 策略模式

import (
	"testing"
)

func TestNewEncryptStrategy(t *testing.T) {
	data := []byte("hello world")
	rsa, _ := NewEncryptStrategy("rsa")
	rsa.Encrypt(data)

	md5, _ := NewEncryptStrategy("md5")
	md5.Encrypt(data)
}
