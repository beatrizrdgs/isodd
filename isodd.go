package isodd

import ie "github.com/beatrizrdgs/iseven"

// IsOdd returns true if the number is odd, false otherwise.
func IsOdd(n int) bool {
	return !ie.IsEven(n)
}
