package utils

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

const (
	// DefaultPageSize is the default number of items per page
	DefaultPageSize = 20
	// MaxPageSize is the maximum allowed items per page
	MaxPageSize = 100
)

// CursorPagination represents cursor-based pagination parameters
type CursorPagination struct {
	Cursor   string `json:"cursor,omitempty"`
	Limit    int    `json:"limit"`
	HasMore  bool   `json:"has_more"`
	NextCursor string `json:"next_cursor,omitempty"`
}

// PageInfo contains pagination metadata
type PageInfo struct {
	HasNextPage     bool   `json:"has_next_page"`
	HasPreviousPage bool   `json:"has_previous_page"`
	StartCursor     string `json:"start_cursor,omitempty"`
	EndCursor       string `json:"end_cursor,omitempty"`
}

// PaginatedResponse is a generic response with pagination
type PaginatedResponse struct {
	Data     interface{} `json:"data"`
	PageInfo PageInfo    `json:"page_info"`
	Total    int64       `json:"total,omitempty"`
}

// CursorData represents the data encoded in a cursor
type CursorData struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

// EncodeCursor encodes cursor data to base64
func EncodeCursor(id uint, createdAt time.Time) (string, error) {
	data := CursorData{
		ID:        id,
		CreatedAt: createdAt,
	}
	
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	
	return base64.StdEncoding.EncodeToString(jsonData), nil
}

// DecodeCursor decodes a base64 cursor to cursor data
func DecodeCursor(cursor string) (*CursorData, error) {
	if cursor == "" {
		return nil, nil
	}
	
	jsonData, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return nil, err
	}
	
	var data CursorData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}
	
	return &data, nil
}

// ValidatePageSize validates and returns a safe page size
func ValidatePageSize(size int) int {
	if size <= 0 {
		return DefaultPageSize
	}
	if size > MaxPageSize {
		return MaxPageSize
	}
	return size
}

// NewPaginatedResponse creates a new paginated response
func NewPaginatedResponse(data interface{}, pageInfo PageInfo, total int64) *PaginatedResponse {
	return &PaginatedResponse{
		Data:     data,
		PageInfo: pageInfo,
		Total:    total,
	}
}
