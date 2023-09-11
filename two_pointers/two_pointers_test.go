package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrom(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{s: "", want: true},
		{s: "abba", want: true},
		{s: "abc", want: false},
		{s: "a b b a", want: true},
		{s: "A b b A", want: true},
		{s: "A b b A", want: true},
		{s: "abca", want: false},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, isPalindrome(test.s))
	}
}
