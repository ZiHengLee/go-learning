package 门面模式

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFacade(t *testing.T) {
	service := UserService{}
	user, err := service.LoginOrRegister(88888888, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}
