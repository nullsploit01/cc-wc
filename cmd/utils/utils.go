package utils

import (
	"bufio"
	"os"
)

func ByteCount(file *os.File) int {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	byteCount := 0
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		byteCount += len(line)
		lineCount += 1
		if scanner.Err() == nil {
			byteCount++ // Add 1 for each newline character except possibly for the last line if it doesn't end with a newline.
		}
	}
	return byteCount + lineCount
}

func LineCount(file *os.File) int {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount
}

func WordCount(file *os.File) int {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		wordCount += 1
	}
	return wordCount
}
