package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

// ByteArrayToString : Convert a []byte array to string with delimator
func ByteArrayToString(a []byte, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

// StringToIntArray : Convert a []byte array to string with delimator
func StringToIntArray(a string, delim string) ([]int, error) {
	var m []int
	s := strings.Split(a, delim)
	for _, r := range s {
		i, err := strconv.Atoi(r)
		if err != nil {
			return nil, err
		}
		m = append(m, i)
	}
	return m, nil
}
