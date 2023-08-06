package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount("John", "Doe", "123456")
	assert.Nil(t, err)

	fmt.Println(acc)
}
