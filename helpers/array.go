package helpers

import "fmt"
import "strings"

func ByteArrayToString(a []byte, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
