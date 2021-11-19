package logger

import (
	"fmt"
	"strings"
)

func Print(format string, args ...interface{}) {
	format = strings.TrimSuffix(format, "\n") + "\n"
	fmt.Printf(format, args...)
}

func Info(format string, args ...interface{}) {
	format = string(Cyan) + "[INFO] " + string(Reset) + strings.TrimSuffix(format, "\n") + "\n"
	fmt.Printf(format, args...)
}

func Error(format string, args ...interface{}) {
	format = string(Red) + "[ERROR] " + string(Reset) + strings.TrimSuffix(format, "\n") + "\n"
	fmt.Printf(format, args...)
}

func Success(format string, args ...interface{}) {
	format = string(Green) + "[SUCCESS] " + string(Reset) + strings.TrimSuffix(format, "\n") + "\n"
	fmt.Printf(format, args...)
}

func Debug(format string, args ...interface{}) {
	format = string(Gray) + "[DEBUG] " + string(Reset) + strings.TrimSuffix(format, "\n") + "\n"
	fmt.Printf(format, args...)
}

func Warn(format string, args ...interface{}) {
	format = string(Yellow) + "[WARNING] " + string(Reset) + strings.TrimSuffix(format, "\n") + "\n"
	fmt.Printf(format, args...)
}

func Panic(format string, args ...interface{}) {
	format = string(Purple) + "[PANIC] " + string(Reset) + strings.TrimSuffix(format, "\n") + "\n"
	fmt.Printf(format, args...)
}

func Custom(suffix string, color Color, format string, args ...interface{}) {
	suffix = strings.TrimSuffix(suffix, " ") + " "
	format = string(color) + suffix + string(Reset) + strings.TrimSuffix(format, "\n") + "\n"
	fmt.Printf(format, args...)
}
