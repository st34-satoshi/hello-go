package sample

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

func WithComment(s string) string{
	return s + "Hello World"
}