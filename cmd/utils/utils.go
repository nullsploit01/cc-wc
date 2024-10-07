package utils

import "strings"

func ByteCount(data string) int {
	return len([]byte(data))
}

func LineCount(data string) int {
	return len(strings.Split(data, "\n"))
}
