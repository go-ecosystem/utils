package log

import (
	"github.com/fatih/color"
)

// E log error
func E(format string, a ...interface{}) {
	color.Red(format, a...)
}

// W log warning
func W(format string, a ...interface{}) {
	color.Yellow(format, a...)
}

// H log hint
func H(format string, a ...interface{}) {
	color.Green(format, a...)
}

// I log info
func I(format string, a ...interface{}) {
	color.Blue(format, a...)
}

// V log verbose
func V(format string, a ...interface{}) {
	color.White(format, a...)
}
