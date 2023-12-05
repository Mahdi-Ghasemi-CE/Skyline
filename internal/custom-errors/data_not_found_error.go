package custom_errors

import (
	"errors"
	"fmt"
)

func CustomDataNotFoundError(dataName string) error {
	text := fmt.Sprintf("The %s is not found !", dataName)
	return errors.New(text)
}

func DataNotFoundError() error {
	text := fmt.Sprintf("The data is not found ")
	return errors.New(text)
}
