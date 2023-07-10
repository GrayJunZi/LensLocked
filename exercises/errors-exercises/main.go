package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("Not Found")
)

type TargetErr struct{}

func (t TargetErr) Error() string {
	return "测试目标错误"
}

func main() {
	err := createNotFound()
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Is ErrNotFound")
	}

	var t TargetErr
	fmt.Println(errors.As(err, &t))
}

func createNotFound() error {
	return ErrNotFound
}
