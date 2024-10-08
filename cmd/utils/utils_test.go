package utils_test

import (
	"os"
	"testing"

	"github.com/nullsploit01/cc-wc/cmd/utils"
)

func createMockFile(data string) *os.File {
	r, w, _ := os.Pipe()
	w.Write([]byte(data))
	w.Close()
	return r
}

func TestByteCount(t *testing.T) {
	data := "Hello, World!\nThis is a test string.\n"
	file := createMockFile(data)
	defer file.Close()

	expected := 37 // byte count of data
	result := utils.ByteCount(file)

	if result != expected {
		t.Errorf("ByteCount was incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestLineCount(t *testing.T) {
	data := "First line\nSecond line\nThird line\n"
	file := createMockFile(data)
	defer file.Close()

	expected := 3 // number of lines
	result := utils.LineCount(file)

	if result != expected {
		t.Errorf("LineCount was incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestWordCount(t *testing.T) {
	data := "Hello, who are you?\nAre you testing this?"
	file := createMockFile(data)
	defer file.Close()

	expected := 8 // number of words
	result := utils.WordCount(file)

	if result != expected {
		t.Errorf("WordCount was incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestCharacterCount(t *testing.T) {
	data := "123456789\n"
	file := createMockFile(data)
	defer file.Close()

	expected := 10 // 9 digits + 1 newline
	result := utils.CharacterCount(file)

	if result != expected {
		t.Errorf("CharacterCount was incorrect, got: %d, want: %d.", result, expected)
	}
}
