package utils

import (
	"reflect"
	"strings"
	"testing"
)

func TestContainsString(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}
	
	tests := []struct {
		name     string
		item     string
		expected bool
	}{
		{"exists", "banana", true},
		{"not exists", "orange", false},
		{"empty string", "", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsString(slice, tt.item)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestUniqueStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"no duplicates", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"with duplicates", []string{"a", "b", "a", "c", "b"}, []string{"a", "b", "c"}},
		{"all duplicates", []string{"a", "a", "a"}, []string{"a"}},
		{"empty slice", []string{}, []string{}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UniqueStrings(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestRemoveString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		item     string
		expected []string
	}{
		{"remove existing", []string{"a", "b", "c"}, "b", []string{"a", "c"}},
		{"remove non-existing", []string{"a", "b", "c"}, "d", []string{"a", "b", "c"}},
		{"remove all occurrences", []string{"a", "b", "a", "c"}, "a", []string{"b", "c"}},
		{"empty slice", []string{}, "a", []string{}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveString(tt.slice, tt.item)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestChunkStrings(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		chunkSize int
		expected  [][]string
	}{
		{"normal chunking", []string{"a", "b", "c", "d", "e"}, 2, 
			[][]string{{"a", "b"}, {"c", "d"}, {"e"}}},
		{"exact division", []string{"a", "b", "c", "d"}, 2,
			[][]string{{"a", "b"}, {"c", "d"}}},
		{"chunk size 1", []string{"a", "b"}, 1,
			[][]string{{"a"}, {"b"}}},
		{"chunk size larger than slice", []string{"a", "b"}, 5,
			[][]string{{"a", "b"}}},
		{"zero chunk size", []string{"a", "b"}, 0,
			[][]string{{"a", "b"}}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ChunkStrings(tt.slice, tt.chunkSize)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFilterStrings(t *testing.T) {
	slice := []string{"apple", "banana", "cherry", "apricot"}
	
	// Filter strings starting with 'a'
	result := FilterStrings(slice, func(s string) bool {
		return len(s) > 0 && s[0] == 'a'
	})
	
	expected := []string{"apple", "apricot"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMapStrings(t *testing.T) {
	slice := []string{"hello", "world"}
	
	// Convert to uppercase
	result := MapStrings(slice, func(s string) string {
		return strings.ToUpper(s)
	})
	
	expected := []string{"HELLO", "WORLD"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
