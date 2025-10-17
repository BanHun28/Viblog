package utils

import (
	"encoding/base64"
	"testing"
	"time"
)

func TestEncodeCursor(t *testing.T) {
	id := uint(123)
	createdAt := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	
	cursor, err := EncodeCursor(id, createdAt)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	
	if cursor == "" {
		t.Error("expected non-empty cursor")
	}
	
	// Verify it's valid base64
	_, err = base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		t.Errorf("cursor is not valid base64: %v", err)
	}
}

func TestDecodeCursor(t *testing.T) {
	id := uint(123)
	createdAt := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	
	cursor, _ := EncodeCursor(id, createdAt)
	
	data, err := DecodeCursor(cursor)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	
	if data == nil {
		t.Fatal("expected non-nil cursor data")
	}
	
	if data.ID != id {
		t.Errorf("expected ID %d, got %d", id, data.ID)
	}
	
	if !data.CreatedAt.Equal(createdAt) {
		t.Errorf("expected time %v, got %v", createdAt, data.CreatedAt)
	}
}

func TestDecodeCursor_Empty(t *testing.T) {
	data, err := DecodeCursor("")
	if err != nil {
		t.Errorf("expected no error for empty cursor, got %v", err)
	}
	
	if data != nil {
		t.Error("expected nil data for empty cursor")
	}
}

func TestDecodeCursor_Invalid(t *testing.T) {
	_, err := DecodeCursor("invalid-base64!!!")
	if err == nil {
		t.Error("expected error for invalid cursor")
	}
}

func TestValidatePageSize(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{"valid size", 10, 10},
		{"default for zero", 0, DefaultPageSize},
		{"default for negative", -5, DefaultPageSize},
		{"max for too large", 200, MaxPageSize},
		{"max page size", MaxPageSize, MaxPageSize},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidatePageSize(tt.size)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestNewPaginatedResponse(t *testing.T) {
	data := []string{"item1", "item2"}
	pageInfo := PageInfo{
		HasNextPage: true,
		EndCursor:   "cursor123",
	}
	total := int64(100)
	
	response := NewPaginatedResponse(data, pageInfo, total)
	
	if response == nil {
		t.Fatal("expected non-nil response")
	}
	
	if response.Total != total {
		t.Errorf("expected total %d, got %d", total, response.Total)
	}
	
	if !response.PageInfo.HasNextPage {
		t.Error("expected HasNextPage to be true")
	}
}
