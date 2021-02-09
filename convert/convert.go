package convert

import (
	"fmt"
	"strings"
)

// IntArrayToString convert an []int to string
func IntArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
