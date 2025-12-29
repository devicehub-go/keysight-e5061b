package utils

import (
	"fmt"
	"strings"
)

/* Join an integer array in a single string */
func JoinIntArray(array []int, delimiter string) string {
	arrayStr := fmt.Sprint(array)
	fields := strings.Fields(arrayStr)
	return strings.Trim(strings.Join(fields, delimiter), "[]")
}
