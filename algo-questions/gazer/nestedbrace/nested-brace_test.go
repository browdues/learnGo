package nestedbrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestIsValid(t *testing.T) {
	cases := []struct {
		name     string
		in       string
		expected bool
	}{
		{"valid nesting - braces", "{}", true},
		{"valid nesting - brackets", "{}", true},
		{"valid nesting - parenthesis", "{}", true},
		{"invalid nesting - braces", "{{}", false},
		{"invalid nesting - braces", "{[}]", false},
		{"invalid nesting - braces", "{({)})", false},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			want := test.expected
			got := NestIsValid(test.in)

			assert.Equal(t, want, got)
		})
	}
}

func TestIsOpener(t *testing.T) {
	cases := []struct {
		name     string
		in       rune
		expected bool
	}{
		{"square bracket", '[', true},
		{"curly brace", '{', true},
		{"paren", '(', true},
		{"random char", 'y', false},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			want := test.expected
			got := isOpener(test.in)
			assert.Equal(t, want, got)
		})
	}
}

func TestIsCloser(t *testing.T) {
	cases := []struct {
		name     string
		in       rune
		expected bool
	}{
		{"square bracket", ']', true},
		{"curly brace", '}', true},
		{"paren", ')', true},
		{"random char", 'y', false},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			want := test.expected
			got := isCloser(test.in)
			assert.Equal(t, want, got)
		})
	}
}
