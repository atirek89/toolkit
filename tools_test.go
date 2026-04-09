package toolkit_test

import (
	"strings"
	"testing"

	"github.com/atirek89/toolkit"
)

func TestTools_RandomString(t *testing.T) {
	var testTools toolkit.Tools

	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Error("wrong length for random string returned")
	} else {
		t.Logf("Test Passed String returned: %s", s)
	}
}

func TestTools_RandomStringNew(t *testing.T) {
	var to toolkit.Tools

	// Test cases for different lengths
	testCases := []struct {
		name string
		n    int
	}{
		{"zero length", 0},
		{"length 1", 1},
		{"length 5", 5},
		{"length 10", 10},
		{"length 100", 100},
	}

	// Allowed characters from the source
	allowedChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := to.RandomString(tc.n)

			// Check length
			if len(got) != tc.n {
				t.Errorf("RandomString(%d) length = %d, want %d", tc.n, len(got), tc.n)
			}

			// Check all characters are allowed
			for _, char := range got {
				if !strings.ContainsRune(allowedChars, char) {
					t.Errorf("RandomString(%d) contains invalid character: %c", tc.n, char)
				}
			}
		})
	}
}

func TestTools_RandomString_Uniqueness(t *testing.T) {
	var to toolkit.Tools

	// Generate multiple strings and check they are likely unique
	// (though not guaranteed, very unlikely to collide)
	strings := make(map[string]bool)
	for i := 0; i < 100; i++ {
		s := to.RandomString(10)
		if strings[s] {
			t.Errorf("RandomString generated duplicate: %s", s)
		}
		strings[s] = true
	}
}
