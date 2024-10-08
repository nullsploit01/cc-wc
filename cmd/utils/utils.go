package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ByteCount(file *os.File) int {
	defer file.Seek(0, 0)

	bufferSize := 32 * 1024 // 32 KB
	buf := make([]byte, bufferSize)
	totalBytes := 0

	for {
		n, err := file.Read(buf)
		totalBytes += n
		if err != nil {
			if err == io.EOF {
				break // We reached the end of the file
			}
			fmt.Printf("Failed to read file: %v\n", err)
			return -1
		}
	}

	return totalBytes
}

func LineCount(file *os.File) int {
	defer file.Seek(0, 0)

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}

func WordCount(file *os.File) int {
	defer file.Seek(0, 0)

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		wordCount += 1
	}

	return wordCount
}

func CharacterCount(file *os.File) int {
	defer file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	charCount := 0

	for scanner.Scan() {
		charCount++
	}

	return charCount
}
