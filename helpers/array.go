package helpers

import "fmt"
import "strings"

// ByteArrayToString : Convert a []byte array to string with delimator
func ByteArrayToString(a []byte, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
