package otp

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Generate creates a new 6-digit OTP
func Generate() (string, error) {
	const otpLength = 6
	const charset = "0123456789"
	
	otp := make([]byte, otpLength)
	charsetLength := big.NewInt(int64(len(charset)))
	
	for i := 0; i < otpLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %w", err)
		}
		otp[i] = charset[randomIndex.Int64()]
	}
	
	return string(otp), nil
}

// Validate checks if the provided OTP matches the expected OTP
func Validate(providedOTP, expectedOTP string) bool {
	if len(providedOTP) != 6 || len(expectedOTP) != 6 {
		return false
	}
	
	// Simple string comparison - in production you might want constant-time comparison
	return providedOTP == expectedOTP
}

// IsValidFormat checks if the OTP has the correct format (6 digits)
func IsValidFormat(otp string) bool {
	if len(otp) != 6 {
		return false
	}
	
	for _, char := range otp {
		if char < '0' || char > '9' {
			return false
		}
	}
	
	return true
}