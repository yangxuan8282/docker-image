package printer

import (
	"fmt"

	"github.com/fatih/color"
)

// Trace traces info
func Trace(format string, values ...interface{}) {
	if EnableLogf {
		Logf(format, values...)
		return
	}
	for i, v := range values {
		values[i] = color.CyanString("%v", v)
	}
	fmt.Printf(format+"\n", values...)
}

// Info traces log in Info level
func Info(format string, values ...interface{}) {
	if EnableLogf {
		Logf(format, values...)
		return
	}
	for i, v := range values {
		values[i] = color.GreenString("%v", v)
	}
	fmt.Printf(format+"\n", values...)
}

// Warn traces log in Warn level
func Warn(format string, values ...interface{}) {
	if EnableLogf {
		Logf(format, values...)
		return
	}
	for i, v := range values {
		values[i] = color.YellowString("%v", v)
	}
	fmt.Printf(format+"\n", values...)
}

// Error traces logs in Error level
func Error(format string, values ...interface{}) {
	if EnableLogf {
		Logf(format, values...)
		return
	}
	for i, v := range values {
		values[i] = color.RedString("%v", v)
	}
	fmt.Printf(format+"\n", values...)
}

// Print traces logs
func Print(format string, values ...interface{}) {
	if EnableLogf {
		Logf(format, values...)
		return
	}
	color.White(format, values...)
}
