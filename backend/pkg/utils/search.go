package utils

import (
	"regexp"
	"strings"
)

// HighlightMatches highlights search query matches in text
func HighlightMatches(text, query string, prefix, suffix string) string {
	if query == "" {
		return text
	}
	
	// Escape special regex characters in query
	escapedQuery := regexp.QuoteMeta(query)
	
	// Create case-insensitive regex
	re := regexp.MustCompile("(?i)" + escapedQuery)
	
	// Replace all matches with highlighted version
	return re.ReplaceAllStringFunc(text, func(match string) string {
		return prefix + match + suffix
	})
}

// HighlightMatchesHTML highlights search matches with HTML tags
func HighlightMatchesHTML(text, query string) string {
	return HighlightMatches(text, query, "<mark>", "</mark>")
}

// ExtractSnippet extracts a snippet around the first match of query
func ExtractSnippet(text, query string, maxLen int) string {
	if query == "" || text == "" {
		return Truncate(text, maxLen, "...")
	}
	
	// Find the first occurrence (case-insensitive)
	lowerText := strings.ToLower(text)
	lowerQuery := strings.ToLower(query)
	
	index := strings.Index(lowerText, lowerQuery)
	if index == -1 {
		// No match found, return truncated text
		return Truncate(text, maxLen, "...")
	}
	
	// Calculate snippet boundaries
	start := index - maxLen/2
	if start < 0 {
		start = 0
	}
	
	end := start + maxLen
	if end > len(text) {
		end = len(text)
	}
	
	// Adjust start if end was capped
	if end-start < maxLen && start > 0 {
		start = end - maxLen
		if start < 0 {
			start = 0
		}
	}
	
	snippet := text[start:end]
	
	// Add ellipsis
	if start > 0 {
		snippet = "..." + snippet
	}
	if end < len(text) {
		snippet = snippet + "..."
	}
	
	return snippet
}

// NormalizeSearchQuery normalizes a search query
func NormalizeSearchQuery(query string) string {
	// Trim whitespace
	query = strings.TrimSpace(query)
	
	// Convert to lowercase for case-insensitive search
	query = strings.ToLower(query)
	
	// Remove multiple spaces
	query = regexp.MustCompile(`\s+`).ReplaceAllString(query, " ")
	
	return query
}

// ParseSearchQuery parses a search query into individual terms
func ParseSearchQuery(query string) []string {
	query = NormalizeSearchQuery(query)
	
	if query == "" {
		return []string{}
	}
	
	// Split by space
	terms := strings.Split(query, " ")
	
	// Filter out empty terms
	var result []string
	for _, term := range terms {
		if term != "" {
			result = append(result, term)
		}
	}
	
	return result
}

// BuildFullTextSearchQuery builds a PostgreSQL full-text search query
func BuildFullTextSearchQuery(query string) string {
	terms := ParseSearchQuery(query)
	
	if len(terms) == 0 {
		return ""
	}
	
	// Join terms with & for AND search
	return strings.Join(terms, " & ")
}

// ContainsAllTerms checks if text contains all search terms
func ContainsAllTerms(text string, terms []string) bool {
	lowerText := strings.ToLower(text)
	
	for _, term := range terms {
		if !strings.Contains(lowerText, strings.ToLower(term)) {
			return false
		}
	}
	
	return true
}
