package logger

import "fmt"

type Logger struct{}

func New(name string) *Logger {
	fmt.Println("Constructing Logger")
	return &Logger{}
}
