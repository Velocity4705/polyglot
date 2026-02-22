package ui

import (
	"fmt"
	"os"
)

// ANSI color codes
const (
	Reset   = "\033[0m"
	Bold    = "\033[1m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Gray    = "\033[90m"
)

var (
	// ColorsEnabled controls whether colors are used
	ColorsEnabled = true
)

func init() {
	// Disable colors if NO_COLOR env var is set or not a TTY
	if os.Getenv("NO_COLOR") != "" {
		ColorsEnabled = false
	}
}

// Colorize wraps text in color codes
func Colorize(color, text string) string {
	if !ColorsEnabled {
		return text
	}
	return color + text + Reset
}

// Success prints success message in green
func Success(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println(Colorize(Green, "✓ "+msg))
}

// Error prints error message in red
func Error(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println(Colorize(Red, "✗ "+msg))
}

// Warning prints warning message in yellow
func Warning(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println(Colorize(Yellow, "⚠ "+msg))
}

// Info prints info message in blue
func Info(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println(Colorize(Blue, "ℹ "+msg))
}

// Step prints step message in cyan
func Step(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println(Colorize(Cyan, "→ "+msg))
}

// Header prints header in bold
func Header(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println(Colorize(Bold, msg))
}

// Dim prints dimmed text in gray
func Dim(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println(Colorize(Gray, msg))
}

// Language prints language name in magenta
func Language(name string) string {
	return Colorize(Magenta+Bold, name)
}

// Command prints command in cyan
func Command(cmd string) string {
	return Colorize(Cyan, cmd)
}

// File prints filename in blue
func File(filename string) string {
	return Colorize(Blue, filename)
}

// Highlight prints highlighted text in yellow
func Highlight(text string) string {
	return Colorize(Yellow+Bold, text)
}
