package utils

import (
	"net/http"
	"strconv"
	"strings"
)

// GetUserAgent extracts user agent from request
func GetUserAgent(r *http.Request) string {
	return r.Header.Get("User-Agent")
}

// GetReferer extracts referer from request
func GetReferer(r *http.Request) string {
	return r.Header.Get("Referer")
}

// IsAJAXRequest checks if the request is an AJAX request
func IsAJAXRequest(r *http.Request) bool {
	return r.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

// GetAcceptLanguage extracts accept language from request
func GetAcceptLanguage(r *http.Request) string {
	return r.Header.Get("Accept-Language")
}

// GetPreferredLanguage extracts the preferred language from Accept-Language header
func GetPreferredLanguage(r *http.Request, supportedLanguages []string) string {
	acceptLang := GetAcceptLanguage(r)
	if acceptLang == "" {
		return "en"
	}
	
	// Parse Accept-Language header (simplified version)
	languages := strings.Split(acceptLang, ",")
	for _, lang := range languages {
		// Remove quality value if present (e.g., "ko;q=0.9")
		lang = strings.TrimSpace(strings.Split(lang, ";")[0])
		
		// Check if it's a supported language
		for _, supported := range supportedLanguages {
			if strings.HasPrefix(lang, supported) {
				return supported
			}
		}
	}
	
	return "en"
}

// SetJSONContentType sets the content type to application/json
func SetJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

// SetHTMLContentType sets the content type to text/html
func SetHTMLContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

// SetNoCacheHeaders sets headers to prevent caching
func SetNoCacheHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

// SetCacheHeaders sets cache headers with max-age
func SetCacheHeaders(w http.ResponseWriter, maxAge int) {
	w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(maxAge))
}

// IsBotRequest checks if the request is from a bot
func IsBotRequest(r *http.Request) bool {
	ua := strings.ToLower(GetUserAgent(r))
	bots := []string{
		"bot", "crawler", "spider", "scraper",
		"googlebot", "bingbot", "slurp", "duckduckbot",
		"baiduspider", "yandexbot", "facebookexternalhit",
	}
	
	for _, bot := range bots {
		if strings.Contains(ua, bot) {
			return true
		}
	}
	
	return false
}
