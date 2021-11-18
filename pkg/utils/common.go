package utils

import "fmt"

// StringFormat - formats a string with the given arguments.
func StringFormat(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
