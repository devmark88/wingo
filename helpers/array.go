package helpers

import "fmt"
import "strings"

func ByteArrayToString(a []byte, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
func ReverseArray(data []*interface{}) {
	d := make([]*interface{}, len(data))
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		d[i], d[j] = data[j], data[i]
	}
	data = d
}
