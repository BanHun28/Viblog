package utils

// MinInt returns the minimum of two integers
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt returns the maximum of two integers
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinUint returns the minimum of two uints
func MinUint(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

// MaxUint returns the maximum of two uints
func MaxUint(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

// ClampInt clamps an integer between min and max values
func ClampInt(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// ClampUint clamps a uint between min and max values
func ClampUint(value, min, max uint) uint {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// AbsInt returns the absolute value of an integer
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// DivideRoundUp performs division and rounds up to the nearest integer
func DivideRoundUp(dividend, divisor int) int {
	if divisor == 0 {
		return 0
	}
	return (dividend + divisor - 1) / divisor
}
