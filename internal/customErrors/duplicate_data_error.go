package customErrors

import (
	"errors"
	"fmt"
)

func DuplicateDataError(dataName string) error {
	text := fmt.Sprintf("this %s is duplicated , please try again with new data ", dataName)
	return errors.New(text)
}
