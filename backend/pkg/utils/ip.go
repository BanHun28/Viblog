package utils

import (
	"net"
	"net/http"
	"strings"
)

// GetIPFromRequest extracts the real IP address from an HTTP request
func GetIPFromRequest(r *http.Request) string {
	// Check X-Forwarded-For header (proxy/load balancer)
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// X-Forwarded-For can contain multiple IPs, take the first one
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			ip := strings.TrimSpace(ips[0])
			if isValidIP(ip) {
				return ip
			}
		}
	}
	
	// Check X-Real-IP header
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		ip := strings.TrimSpace(xri)
		if isValidIP(ip) {
			return ip
		}
	}
	
	// Fallback to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// If SplitHostPort fails, use RemoteAddr as-is
		return r.RemoteAddr
	}
	
	return ip
}

// isValidIP checks if a string is a valid IP address
func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsLocalIP checks if an IP is a local/private IP
func IsLocalIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	
	// Check if it's loopback
	if parsedIP.IsLoopback() {
		return true
	}
	
	// Check if it's private
	if parsedIP.IsPrivate() {
		return true
	}
	
	return false
}

// NormalizeIP normalizes an IP address (e.g., converts IPv4-mapped IPv6 to IPv4)
func NormalizeIP(ip string) string {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return ip
	}
	
	// Convert IPv4-mapped IPv6 to IPv4
	if ipv4 := parsedIP.To4(); ipv4 != nil {
		return ipv4.String()
	}
	
	return parsedIP.String()
}

// HashIP creates a hashed version of IP for privacy (simple anonymization)
func HashIP(ip string) string {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return ip
	}
	
	// For IPv4, zero out last octet
	if ipv4 := parsedIP.To4(); ipv4 != nil {
		ipv4[3] = 0
		return ipv4.String()
	}
	
	// For IPv6, zero out last 80 bits
	ipv6 := parsedIP.To16()
	for i := 6; i < 16; i++ {
		ipv6[i] = 0
	}
	return ipv6.String()
}
