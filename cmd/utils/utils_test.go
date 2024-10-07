package utils_test

import (
	"testing"

	"github.com/nullsploit01/cc-wc/cmd/utils"
)

func TestByteCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"empty string", "", 0},
		{"ascii string", "hello", 5},
		{"unicode string", "こんにちは", 15}, // Each Japanese character is 3 bytes in UTF-8.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.ByteCount(tt.input)
			if result != tt.expected {
				t.Errorf("ByteCount(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestLineCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"empty string", "", 1},
		{"one line", "hello", 1},
		{"two lines unix", "hello\nworld", 2},
		{"two lines windows", "hello\r\nworld", 2},
		{"multiple lines", "line1\nline2\nline3", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.LineCount(tt.input)
			if result != tt.expected {
				t.Errorf("LineCount(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
