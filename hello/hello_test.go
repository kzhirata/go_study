package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
    assert.Equal(t, getHello(), "Hello World")
}
