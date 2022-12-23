package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSandboxOK(t *testing.T) {
	left := "This is a sandbox project."
	right := "This is a sandbox project."

	assert.Equal(t, left, right)
}

func TestSandboxErr(t *testing.T) {
	t.Skip("Skipp error test")
	t.Error("I am an error!")
}
