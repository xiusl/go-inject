package error

import (
	"errors"
	"fmt"
)

type AppError error

func NewAppError(msg string) AppError {
	return errors.New(msg)
}

func NewAppNameDuplicatedErr(name string) AppError {
	msg := fmt.Sprintf("%s 已经存在", name)
	return NewAppError(msg)
}

func NewAppNotFoundErr(item string) AppError {
	msg := fmt.Sprintf("%s 没有找到", item)
	return NewAppError(msg)
}
