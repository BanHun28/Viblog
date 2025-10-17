package validator

import (
	"net/url"
	"regexp"
	"strings"
	"unicode"
)

var (
	// emailRegex is a basic email validation regex
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

// IsValidEmail validates email format
func IsValidEmail(email string) bool {
	if len(email) == 0 || len(email) > 254 {
		return false
	}
	return emailRegex.MatchString(email)
}

// IsValidURL validates URL format
func IsValidURL(urlStr string) bool {
	if len(urlStr) == 0 {
		return false
	}
	
	u, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return false
	}
	
	// Check if scheme is http or https
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	
	// Check if host is present
	if u.Host == "" {
		return false
	}
	
	return true
}

// PasswordValidationResult contains password validation details
type PasswordValidationResult struct {
	Valid          bool
	MinLength      bool
	HasLetter      bool
	HasNumber      bool
	HasSpecialChar bool
}

// ValidatePassword validates password against policy (min 8 chars, letters+numbers+special chars)
func ValidatePassword(password string) PasswordValidationResult {
	result := PasswordValidationResult{
		MinLength:      len(password) >= 8,
		HasLetter:      false,
		HasNumber:      false,
		HasSpecialChar: false,
	}
	
	for _, char := range password {
		switch {
		case unicode.IsLetter(char):
			result.HasLetter = true
		case unicode.IsNumber(char):
			result.HasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			result.HasSpecialChar = true
		}
	}
	
	result.Valid = result.MinLength && result.HasLetter && result.HasNumber && result.HasSpecialChar
	return result
}

// IsValidPassword checks if password meets requirements (simplified version)
func IsValidPassword(password string) bool {
	result := ValidatePassword(password)
	return result.Valid
}

// IsValidNickname validates nickname (alphanumeric, underscore, hyphen, 2-20 chars)
func IsValidNickname(nickname string) bool {
	if len(nickname) < 2 || len(nickname) > 20 {
		return false
	}
	
	nicknameRegex := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	return nicknameRegex.MatchString(nickname)
}

// IsValidTagName validates tag name (no special chars except hyphen, 1-30 chars)
func IsValidTagName(tag string) bool {
	tag = strings.TrimSpace(tag)
	if len(tag) < 1 || len(tag) > 30 {
		return false
	}
	
	tagRegex := regexp.MustCompile(`^[a-zA-Z0-9가-힣-]+$`)
	return tagRegex.MatchString(tag)
}

// IsValidCategorySlug validates category slug (lowercase, hyphen, 1-50 chars)
func IsValidCategorySlug(slug string) bool {
	if len(slug) < 1 || len(slug) > 50 {
		return false
	}
	
	slugRegex := regexp.MustCompile(`^[a-z0-9-]+$`)
	return slugRegex.MatchString(slug)
}

// SanitizeString removes leading/trailing whitespace and normalizes internal whitespace
func SanitizeString(s string) string {
	s = strings.TrimSpace(s)
	// Replace multiple spaces with single space
	spaceRegex := regexp.MustCompile(`\s+`)
	return spaceRegex.ReplaceAllString(s, " ")
}
