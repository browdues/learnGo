package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Shideh")

	got := buffer.String()
	want := "Hi Shideh"

	assert.Equal(t, want, got)
}
