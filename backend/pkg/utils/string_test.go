package utils

import (
	"testing"
)

func TestSlugify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple text", "Hello World", "hello-world"},
		{"with special chars", "Hello, World!", "hello-world"},
		{"with underscores", "hello_world_test", "hello-world-test"},
		{"multiple spaces", "hello   world", "hello-world"},
		{"leading/trailing hyphens", "  -hello-world-  ", "hello-world"},
		{"consecutive hyphens", "hello---world", "hello-world"},
		{"numbers", "test-123-abc", "test-123-abc"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Slugify(tt.input)
			if result != tt.expected {
				t.Errorf("Slugify(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxLen   int
		suffix   string
		expected string
	}{
		{"no truncation needed", "hello", 10, "...", "hello"},
		{"exact length", "hello", 5, "...", "hello"},
		{"truncate with suffix", "hello world", 8, "...", "hello..."},
		{"suffix longer than maxLen", "hello world", 2, "...", "he"},
		{"empty suffix", "hello world", 5, "", "hello"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Truncate(tt.input, tt.maxLen, tt.suffix)
			if result != tt.expected {
				t.Errorf("Truncate(%q, %d, %q) = %q, expected %q", 
					tt.input, tt.maxLen, tt.suffix, result, tt.expected)
			}
		})
	}
}

func TestStripHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"no HTML", "plain text", "plain text"},
		{"simple tags", "<p>hello</p>", "hello"},
		{"nested tags", "<div><p>hello</p></div>", "hello"},
		{"with entities", "hello&nbsp;world&lt;test&gt;", "hello world<test>"},
		{"multiple tags", "<strong>bold</strong> and <em>italic</em>", "bold and italic"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StripHTML(tt.input)
			if result != tt.expected {
				t.Errorf("StripHTML(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"spaces only", "   ", true},
		{"tabs only", "\t\t", true},
		{"newlines only", "\n\n", true},
		{"mixed whitespace", " \t\n ", true},
		{"has content", "hello", false},
		{"content with spaces", "  hello  ", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsBlank(tt.input)
			if result != tt.expected {
				t.Errorf("IsBlank(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		name       string
		str        string
		substrings []string
		expected   bool
	}{
		{"contains one", "hello world", []string{"world", "foo"}, true},
		{"contains multiple", "hello world", []string{"hello", "world"}, true},
		{"contains none", "hello world", []string{"foo", "bar"}, false},
		{"empty substrings", "hello", []string{}, false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsAny(tt.str, tt.substrings)
			if result != tt.expected {
				t.Errorf("ContainsAny(%q, %v) = %v, expected %v", 
					tt.str, tt.substrings, result, tt.expected)
			}
		})
	}
}
