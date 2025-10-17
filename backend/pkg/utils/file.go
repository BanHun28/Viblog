package utils

import (
	"mime"
	"path/filepath"
	"strings"
)

// GetMIMEType returns the MIME type for a given file extension
func GetMIMEType(filename string) string {
	ext := filepath.Ext(filename)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		return "application/octet-stream"
	}
	return mimeType
}

// GetFileExtension returns the file extension from a filename
func GetFileExtension(filename string) string {
	ext := filepath.Ext(filename)
	return strings.TrimPrefix(ext, ".")
}

// IsImageFile checks if a file is an image based on extension
func IsImageFile(filename string) bool {
	ext := strings.ToLower(GetFileExtension(filename))
	imageExts := []string{"jpg", "jpeg", "png", "gif", "bmp", "svg", "webp"}
	
	for _, imgExt := range imageExts {
		if ext == imgExt {
			return true
		}
	}
	
	return false
}

// IsMarkdownFile checks if a file is a markdown file
func IsMarkdownFile(filename string) bool {
	ext := strings.ToLower(GetFileExtension(filename))
	return ext == "md" || ext == "markdown"
}

// SanitizeFilename removes potentially dangerous characters from filename
func SanitizeFilename(filename string) string {
	// Remove path separators
	filename = strings.ReplaceAll(filename, "/", "")
	filename = strings.ReplaceAll(filename, "\\", "")
	
	// Remove null bytes
	filename = strings.ReplaceAll(filename, "\x00", "")
	
	// Trim spaces and dots
	filename = strings.TrimSpace(filename)
	filename = strings.Trim(filename, ".")
	
	return filename
}
