package assert

import (
	"fmt"

	"github.com/gookit/color"
)

// formatStr formats string by argument placing
func formatStr(f string, a ...interface{}) string {
	return fmt.Sprintf(f, a...)
}

// red prints text in red
func red(format string, args ...interface{}) string {
	return color.Red.Text(fmt.Sprintf("%s", formatStr(format, args...)))
}

// green prints text in green
func green(format string, args ...interface{}) string {
	return color.Green.Text(fmt.Sprintf("%s", formatStr(format, args...)))
}
