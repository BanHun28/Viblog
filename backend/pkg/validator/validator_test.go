package validator

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"valid email", "test@example.com", true},
		{"valid with subdomain", "user@mail.example.com", true},
		{"valid with plus", "user+tag@example.com", true},
		{"invalid no @", "testexample.com", false},
		{"invalid no domain", "test@", false},
		{"invalid no user", "@example.com", false},
		{"invalid spaces", "test @example.com", false},
		{"empty string", "", false},
		{"too long", string(make([]byte, 255)) + "@test.com", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidEmail(tt.email)
			if result != tt.expected {
				t.Errorf("IsValidEmail(%q) = %v, expected %v", tt.email, result, tt.expected)
			}
		})
	}
}

func TestIsValidURL(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected bool
	}{
		{"valid http", "http://example.com", true},
		{"valid https", "https://example.com", true},
		{"valid with path", "https://example.com/path/to/resource", true},
		{"valid with query", "https://example.com?key=value", true},
		{"invalid no scheme", "example.com", false},
		{"invalid ftp scheme", "ftp://example.com", false},
		{"invalid no host", "https://", false},
		{"empty string", "", false},
		{"invalid characters", "https://exam ple.com", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidURL(tt.url)
			if result != tt.expected {
				t.Errorf("IsValidURL(%q) = %v, expected %v", tt.url, result, tt.expected)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name           string
		password       string
		expectedValid  bool
		expectedMin    bool
		expectedLetter bool
		expectedNumber bool
		expectedSpecial bool
	}{
		{
			name:            "valid password",
			password:        "Test123!@#",
			expectedValid:   true,
			expectedMin:     true,
			expectedLetter:  true,
			expectedNumber:  true,
			expectedSpecial: true,
		},
		{
			name:            "too short",
			password:        "Te1!",
			expectedValid:   false,
			expectedMin:     false,
			expectedLetter:  true,
			expectedNumber:  true,
			expectedSpecial: true,
		},
		{
			name:            "no letter",
			password:        "12345678!@#",
			expectedValid:   false,
			expectedMin:     true,
			expectedLetter:  false,
			expectedNumber:  true,
			expectedSpecial: true,
		},
		{
			name:            "no number",
			password:        "TestTest!@#",
			expectedValid:   false,
			expectedMin:     true,
			expectedLetter:  true,
			expectedNumber:  false,
			expectedSpecial: true,
		},
		{
			name:            "no special char",
			password:        "TestTest123",
			expectedValid:   false,
			expectedMin:     true,
			expectedLetter:  true,
			expectedNumber:  true,
			expectedSpecial: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidatePassword(tt.password)
			
			if result.Valid != tt.expectedValid {
				t.Errorf("Valid = %v, expected %v", result.Valid, tt.expectedValid)
			}
			if result.MinLength != tt.expectedMin {
				t.Errorf("MinLength = %v, expected %v", result.MinLength, tt.expectedMin)
			}
			if result.HasLetter != tt.expectedLetter {
				t.Errorf("HasLetter = %v, expected %v", result.HasLetter, tt.expectedLetter)
			}
			if result.HasNumber != tt.expectedNumber {
				t.Errorf("HasNumber = %v, expected %v", result.HasNumber, tt.expectedNumber)
			}
			if result.HasSpecialChar != tt.expectedSpecial {
				t.Errorf("HasSpecialChar = %v, expected %v", result.HasSpecialChar, tt.expectedSpecial)
			}
		})
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		expected bool
	}{
		{"valid password", "Test123!@#", true},
		{"too short", "Te1!", false},
		{"no special char", "TestTest123", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidPassword(tt.password)
			if result != tt.expected {
				t.Errorf("IsValidPassword(%q) = %v, expected %v", tt.password, result, tt.expected)
			}
		})
	}
}

func TestIsValidNickname(t *testing.T) {
	tests := []struct {
		name     string
		nickname string
		expected bool
	}{
		{"valid alphanumeric", "user123", true},
		{"valid with underscore", "user_name", true},
		{"valid with hyphen", "user-name", true},
		{"too short", "u", false},
		{"too long", "thisusernameiswaytoolongforvalidation", false},
		{"with spaces", "user name", false},
		{"with special chars", "user@name", false},
		{"empty", "", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidNickname(tt.nickname)
			if result != tt.expected {
				t.Errorf("IsValidNickname(%q) = %v, expected %v", tt.nickname, result, tt.expected)
			}
		})
	}
}

func TestIsValidTagName(t *testing.T) {
	tests := []struct {
		name     string
		tag      string
		expected bool
	}{
		{"valid english", "golang", true},
		{"valid korean", "프로그래밍", true},
		{"valid with hyphen", "web-dev", true},
		{"valid with number", "tech2024", true},
		{"too long", string(make([]rune, 31)), false},
		{"empty", "", false},
		{"with special chars", "tag@name", false},
		{"with spaces trimmed", " tag ", true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidTagName(tt.tag)
			if result != tt.expected {
				t.Errorf("IsValidTagName(%q) = %v, expected %v", tt.tag, result, tt.expected)
			}
		})
	}
}

func TestIsValidCategorySlug(t *testing.T) {
	tests := []struct {
		name     string
		slug     string
		expected bool
	}{
		{"valid lowercase", "web-development", true},
		{"valid with numbers", "tech-2024", true},
		{"uppercase not allowed", "WEB-DEV", false},
		{"spaces not allowed", "web dev", false},
		{"too long", string(make([]byte, 51)), false},
		{"empty", "", false},
		{"special chars", "web_dev", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidCategorySlug(tt.slug)
			if result != tt.expected {
				t.Errorf("IsValidCategorySlug(%q) = %v, expected %v", tt.slug, result, tt.expected)
			}
		})
	}
}

func TestSanitizeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"trim spaces", "  test  ", "test"},
		{"normalize spaces", "test  multiple   spaces", "test multiple spaces"},
		{"tabs to single space", "test\t\ttabs", "test tabs"},
		{"newlines to single space", "test\n\nnewlines", "test newlines"},
		{"mixed whitespace", "  test \t multiple \n spaces  ", "test multiple spaces"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeString(tt.input)
			if result != tt.expected {
				t.Errorf("SanitizeString(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}
