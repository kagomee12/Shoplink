package pkg

import (
	"errors"
	"fmt"
	"shoplink/app/constant"
)

func PanicException_(key string, messages string) {
	errors := errors.New(messages)
	errors = fmt.Errorf("%s: %w", key, errors)
	if errors != nil {
		panic(errors)
	}
}

func PanicException(constant constant.ResponseStatus) {
	PanicException_(constant.GetResponseStatus(), constant.GetResponseMessage())
}
