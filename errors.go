package main

import "errors"

// NewError 创建异常
func NewError(msg string) error {
	return errors.New(msg)
}
