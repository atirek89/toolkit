// Package toolkit provides utility functions for common tasks.
// It includes tools for generating random strings and other helper methods.
package toolkit

import "crypto/rand"

const randStrSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type used to instantiate this module. Any variable of this type will
// have access to all the module with the receiver *Tools
type Tools struct{}

// RandomString generates a random string of length n using cryptographically secure random bytes.
// It uses a predefined set of characters including letters, numbers, and symbols.
// The function ensures randomness by selecting characters based on prime numbers generated from crypto/rand.
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randStrSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))

		s[i] = r[x%y]
	}
	return string(s)
}
