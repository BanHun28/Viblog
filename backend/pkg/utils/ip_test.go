package utils

import (
	"net/http/httptest"
	"testing"
)

func TestGetIPFromRequest(t *testing.T) {
	tests := []struct {
		name            string
		remoteAddr      string
		xForwardedFor   string
		xRealIP         string
		expectedContain string
	}{
		{
			name:            "from RemoteAddr",
			remoteAddr:      "192.168.1.1:12345",
			expectedContain: "192.168.1.1",
		},
		{
			name:            "from X-Forwarded-For",
			remoteAddr:      "127.0.0.1:12345",
			xForwardedFor:   "203.0.113.1, 198.51.100.1",
			expectedContain: "203.0.113.1",
		},
		{
			name:            "from X-Real-IP",
			remoteAddr:      "127.0.0.1:12345",
			xRealIP:         "203.0.113.2",
			expectedContain: "203.0.113.2",
		},
		{
			name:            "X-Forwarded-For priority over X-Real-IP",
			remoteAddr:      "127.0.0.1:12345",
			xForwardedFor:   "203.0.113.1",
			xRealIP:         "203.0.113.2",
			expectedContain: "203.0.113.1",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			req.RemoteAddr = tt.remoteAddr
			
			if tt.xForwardedFor != "" {
				req.Header.Set("X-Forwarded-For", tt.xForwardedFor)
			}
			if tt.xRealIP != "" {
				req.Header.Set("X-Real-IP", tt.xRealIP)
			}
			
			ip := GetIPFromRequest(req)
			
			if ip != tt.expectedContain {
				t.Errorf("expected IP containing %s, got %s", tt.expectedContain, ip)
			}
		})
	}
}

func TestIsLocalIP(t *testing.T) {
	tests := []struct {
		name     string
		ip       string
		expected bool
	}{
		{"localhost", "127.0.0.1", true},
		{"localhost IPv6", "::1", true},
		{"private 192.168", "192.168.1.1", true},
		{"private 10", "10.0.0.1", true},
		{"private 172.16", "172.16.0.1", true},
		{"public IP", "8.8.8.8", false},
		{"public IP 2", "1.1.1.1", false},
		{"invalid IP", "invalid", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsLocalIP(tt.ip)
			if result != tt.expected {
				t.Errorf("IsLocalIP(%s) = %v, expected %v", tt.ip, result, tt.expected)
			}
		})
	}
}

func TestNormalizeIP(t *testing.T) {
	tests := []struct {
		name     string
		ip       string
		expected string
	}{
		{"IPv4", "192.168.1.1", "192.168.1.1"},
		{"IPv6", "2001:db8::1", "2001:db8::1"},
		{"IPv4-mapped IPv6", "::ffff:192.168.1.1", "192.168.1.1"},
		{"invalid IP", "invalid", "invalid"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeIP(tt.ip)
			if result != tt.expected {
				t.Errorf("NormalizeIP(%s) = %s, expected %s", tt.ip, result, tt.expected)
			}
		})
	}
}

func TestHashIP(t *testing.T) {
	tests := []struct {
		name        string
		ip          string
		expectedEnd string
	}{
		{"IPv4", "192.168.1.123", "192.168.1.0"},
		{"localhost", "127.0.0.1", "127.0.0.0"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HashIP(tt.ip)
			if result != tt.expectedEnd {
				t.Errorf("HashIP(%s) = %s, expected %s", tt.ip, result, tt.expectedEnd)
			}
		})
	}
}

func TestHashIP_Privacy(t *testing.T) {
	// Test that hashing removes identifying information
	ip1 := "192.168.1.100"
	ip2 := "192.168.1.200"
	
	hash1 := HashIP(ip1)
	hash2 := HashIP(ip2)
	
	// Both should hash to same value (last octet zeroed)
	if hash1 != hash2 {
		t.Errorf("IPs from same subnet should hash to same value: %s vs %s", hash1, hash2)
	}
	
	expected := "192.168.1.0"
	if hash1 != expected {
		t.Errorf("expected %s, got %s", expected, hash1)
	}
}
