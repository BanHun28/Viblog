package utils

import (
	"regexp"
	"strings"
	"unicode"
)

// Slugify converts a string to a URL-friendly slug
func Slugify(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)
	
	// Replace spaces and underscores with hyphens
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "_", "-")
	
	// Remove all non-alphanumeric characters except hyphens
	reg := regexp.MustCompile(`[^a-z0-9-]+`)
	s = reg.ReplaceAllString(s, "")
	
	// Remove consecutive hyphens
	reg = regexp.MustCompile(`-+`)
	s = reg.ReplaceAllString(s, "-")
	
	// Trim hyphens from start and end
	s = strings.Trim(s, "-")
	
	return s
}

// Truncate truncates a string to the specified length and adds suffix if truncated
func Truncate(s string, maxLen int, suffix string) string {
	if len(s) <= maxLen {
		return s
	}
	
	if len(suffix) >= maxLen {
		return s[:maxLen]
	}
	
	return s[:maxLen-len(suffix)] + suffix
}

// StripHTML removes HTML tags from a string (basic implementation)
func StripHTML(s string) string {
	// Remove all HTML tags
	reg := regexp.MustCompile(`<[^>]*>`)
	s = reg.ReplaceAllString(s, "")
	
	// Replace HTML entities
	s = strings.ReplaceAll(s, "&nbsp;", " ")
	s = strings.ReplaceAll(s, "&lt;", "<")
	s = strings.ReplaceAll(s, "&gt;", ">")
	s = strings.ReplaceAll(s, "&amp;", "&")
	s = strings.ReplaceAll(s, "&quot;", "\"")
	
	return s
}

// SanitizeHTML removes potentially dangerous HTML/script tags
func SanitizeHTML(s string) string {
	// Remove script tags
	scriptReg := regexp.MustCompile(`(?i)<script[^>]*>.*?</script>`)
	s = scriptReg.ReplaceAllString(s, "")
	
	// Remove on* event handlers
	eventReg := regexp.MustCompile(`(?i)\s*on\w+\s*=\s*["'][^"']*["']`)
	s = eventReg.ReplaceAllString(s, "")
	
	// Remove javascript: links
	jsReg := regexp.MustCompile(`(?i)javascript:`)
	s = jsReg.ReplaceAllString(s, "")
	
	return s
}

// ExtractExcerpt extracts a text excerpt from markdown/HTML content
func ExtractExcerpt(content string, maxLen int) string {
	// Strip HTML first
	content = StripHTML(content)
	
	// Remove markdown symbols
	content = strings.ReplaceAll(content, "#", "")
	content = strings.ReplaceAll(content, "*", "")
	content = strings.ReplaceAll(content, "_", "")
	content = strings.ReplaceAll(content, "`", "")
	
	// Normalize whitespace
	content = regexp.MustCompile(`\s+`).ReplaceAllString(content, " ")
	content = strings.TrimSpace(content)
	
	// Truncate to maxLen
	if len(content) <= maxLen {
		return content
	}
	
	// Try to break at word boundary
	lastSpace := strings.LastIndex(content[:maxLen], " ")
	if lastSpace > maxLen/2 {
		return content[:lastSpace] + "..."
	}
	
	return content[:maxLen] + "..."
}

// IsBlank checks if a string is empty or contains only whitespace
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// ContainsAny checks if string contains any of the substrings
func ContainsAny(s string, substrings []string) bool {
	for _, substr := range substrings {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// HasOnlyPrintable checks if string contains only printable characters
func HasOnlyPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
