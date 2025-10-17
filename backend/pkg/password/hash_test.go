package password

import (
	"strings"
	"testing"
)

func TestHash(t *testing.T) {
	password := "testPassword123!@#"
	
	hash, err := Hash(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	
	if hash == "" {
		t.Error("expected hash to be non-empty")
	}
	
	if hash == password {
		t.Error("hash should not equal plain password")
	}
	
	// bcrypt hashes start with $2a$ or $2b$
	if !strings.HasPrefix(hash, "$2") {
		t.Errorf("expected bcrypt hash format, got %s", hash)
	}
}

func TestVerify(t *testing.T) {
	password := "testPassword123!@#"
	hash, _ := Hash(password)
	
	tests := []struct {
		name     string
		password string
		hash     string
		expected bool
	}{
		{
			name:     "correct password",
			password: password,
			hash:     hash,
			expected: true,
		},
		{
			name:     "incorrect password",
			password: "wrongPassword",
			hash:     hash,
			expected: false,
		},
		{
			name:     "empty password",
			password: "",
			hash:     hash,
			expected: false,
		},
		{
			name:     "invalid hash",
			password: password,
			hash:     "invalid_hash",
			expected: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Verify(tt.password, tt.hash)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestHashWithCost(t *testing.T) {
	password := "testPassword123!@#"
	
	tests := []struct {
		name string
		cost int
		err  bool
	}{
		{
			name: "minimum cost",
			cost: 4,
			err:  false,
		},
		{
			name: "default cost",
			cost: DefaultCost,
			err:  false,
		},
		{
			name: "high cost",
			cost: 12,
			err:  false,
		},
		{
			name: "too high cost",
			cost: 32,
			err:  true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashWithCost(password, tt.cost)
			
			if tt.err {
				if err == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if hash == "" {
					t.Error("expected hash to be non-empty")
				}
				
				// Verify the hash works
				if !Verify(password, hash) {
					t.Error("hash verification failed")
				}
			}
		})
	}
}

func TestHashDifferentPasswords(t *testing.T) {
	password1 := "password1"
	password2 := "password2"
	
	hash1, _ := Hash(password1)
	hash2, _ := Hash(password2)
	
	if hash1 == hash2 {
		t.Error("different passwords should produce different hashes")
	}
}

func TestHashSamePasswordTwice(t *testing.T) {
	password := "testPassword123!@#"
	
	hash1, _ := Hash(password)
	hash2, _ := Hash(password)
	
	// bcrypt includes salt, so same password should produce different hashes
	if hash1 == hash2 {
		t.Error("same password hashed twice should produce different hashes (due to salt)")
	}
	
	// But both should verify correctly
	if !Verify(password, hash1) {
		t.Error("first hash verification failed")
	}
	if !Verify(password, hash2) {
		t.Error("second hash verification failed")
	}
}
