package utils

import (
	"strings"
	"testing"
)

func TestHighlightMatches(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		query    string
		prefix   string
		suffix   string
		expected string
	}{
		{
			name:     "simple match",
			text:     "hello world",
			query:    "world",
			prefix:   "[",
			suffix:   "]",
			expected: "hello [world]",
		},
		{
			name:     "case insensitive",
			text:     "Hello World",
			query:    "world",
			prefix:   "[",
			suffix:   "]",
			expected: "Hello [World]",
		},
		{
			name:     "multiple matches",
			text:     "test test test",
			query:    "test",
			prefix:   "<b>",
			suffix:   "</b>",
			expected: "<b>test</b> <b>test</b> <b>test</b>",
		},
		{
			name:     "no match",
			text:     "hello world",
			query:    "foo",
			prefix:   "[",
			suffix:   "]",
			expected: "hello world",
		},
		{
			name:     "empty query",
			text:     "hello world",
			query:    "",
			prefix:   "[",
			suffix:   "]",
			expected: "hello world",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HighlightMatches(tt.text, tt.query, tt.prefix, tt.suffix)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestHighlightMatchesHTML(t *testing.T) {
	text := "hello world"
	query := "world"
	
	result := HighlightMatchesHTML(text, query)
	expected := "hello <mark>world</mark>"
	
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestExtractSnippet(t *testing.T) {
	text := "The quick brown fox jumps over the lazy dog"
	query := "fox"
	maxLen := 30
	
	result := ExtractSnippet(text, query, maxLen)
	
	// Should contain the query
	if !strings.Contains(result, "fox") {
		t.Errorf("snippet should contain query term: %q", result)
	}
	
	// Should have ellipsis (match is in middle)
	if !strings.Contains(result, "...") {
		t.Errorf("snippet should have ellipsis: %q", result)
	}
	
	// Test match at start
	text2 := "Hello world, this is a test"
	query2 := "Hello"
	result2 := ExtractSnippet(text2, query2, maxLen)
	
	if !strings.Contains(result2, "Hello") {
		t.Errorf("snippet should contain query: %q", result2)
	}
	
	// Test no match
	text3 := "Hello world"
	query3 := "xyz"
	result3 := ExtractSnippet(text3, query3, 8)
	
	// Should truncate text
	if len(result3) > 11 { // 8 + "..."
		t.Errorf("snippet should be truncated: %q (len: %d)", result3, len(result3))
	}
}

func TestNormalizeSearchQuery(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected string
	}{
		{"trim spaces", "  hello world  ", "hello world"},
		{"lowercase", "HELLO World", "hello world"},
		{"multiple spaces", "hello   world", "hello world"},
		{"mixed", "  HELLO   World  ", "hello world"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeSearchQuery(tt.query)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestParseSearchQuery(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected []string
	}{
		{"single term", "hello", []string{"hello"}},
		{"multiple terms", "hello world", []string{"hello", "world"}},
		{"with spaces", "  hello   world  ", []string{"hello", "world"}},
		{"empty", "", []string{}},
		{"only spaces", "   ", []string{}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseSearchQuery(tt.query)
			
			if len(result) != len(tt.expected) {
				t.Errorf("expected %d terms, got %d", len(tt.expected), len(result))
				return
			}
			
			for i, term := range result {
				if term != tt.expected[i] {
					t.Errorf("term %d: expected %q, got %q", i, tt.expected[i], term)
				}
			}
		})
	}
}

func TestBuildFullTextSearchQuery(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected string
	}{
		{"single term", "hello", "hello"},
		{"multiple terms", "hello world", "hello & world"},
		{"with normalization", "HELLO  World", "hello & world"},
		{"empty", "", ""},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildFullTextSearchQuery(tt.query)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestContainsAllTerms(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		terms    []string
		expected bool
	}{
		{
			name:     "all terms present",
			text:     "the quick brown fox",
			terms:    []string{"quick", "fox"},
			expected: true,
		},
		{
			name:     "case insensitive",
			text:     "The Quick Brown Fox",
			terms:    []string{"quick", "fox"},
			expected: true,
		},
		{
			name:     "missing term",
			text:     "the quick brown fox",
			terms:    []string{"quick", "dog"},
			expected: false,
		},
		{
			name:     "empty terms",
			text:     "hello world",
			terms:    []string{},
			expected: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsAllTerms(tt.text, tt.terms)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestHighlightMatches_SpecialCharacters(t *testing.T) {
	// Test that special regex characters are properly escaped
	text := "price $100 (approx.)"
	query := "$100"
	
	result := HighlightMatches(text, query, "[", "]")
	expected := "price [$100] (approx.)"
	
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestExtractSnippet_CaseInsensitive(t *testing.T) {
	text := "The quick BROWN fox jumps over the lazy dog"
	query := "brown"
	maxLen := 30
	
	result := ExtractSnippet(text, query, maxLen)
	
	// Should find match regardless of case
	if !strings.Contains(result, "BROWN") {
		t.Errorf("snippet should contain matched term: %q", result)
	}
}
