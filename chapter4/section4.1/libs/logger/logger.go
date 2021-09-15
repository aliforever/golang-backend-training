package logger

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	output io.Writer
}

func NewLogger() *Logger {
	return &Logger{output: os.Stdout}
}

func (l *Logger) SetOutput(writer io.Writer) {
	l.output = writer
}

func (l *Logger) Begin() {
	fmt.Fprintln(l.output, "BEGIN")
}

func (l *Logger) End() {
	fmt.Fprintln(l.output, "END")
}

func (l *Logger) Log(title, message string) {
	fmt.Fprintln(l.output, fmt.Sprintf("%s %s", title, message))
}
