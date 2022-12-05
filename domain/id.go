package domain

import (
	"errors"
	"strings"
)

type ID string

func (id ID) EnsureValid() error {
	if strings.TrimSpace(string(id)) == "" {
		return errors.New("id cannot be empty")
	}

	return nil
}
func (id ID) toInt() int {

	return 1
	//uuid lo crea el front
}
