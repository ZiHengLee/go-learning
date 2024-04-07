package 策略模式

import (
	"fmt"
)

type EncryptStrategy interface {
	Encrypt(data []byte) error
}

var encrypts = map[string]EncryptStrategy{
	"rsa": &rasStg{},
	"md5": &md5Stg{},
}

func NewEncryptStrategy(t string) (EncryptStrategy, error) {
	s, ok := encrypts[t]
	if !ok {
		return nil, fmt.Errorf("not found EncryptStrategy: %s", t)
	}

	return s, nil
}

type rasStg struct{}

func (s *rasStg) Encrypt(data []byte) error {
	fmt.Println("rsa 加密算法")
	return nil
}

type md5Stg struct{}

func (s *md5Stg) Encrypt(data []byte) error {
	fmt.Println("md5 加密算法")
	return nil
}
