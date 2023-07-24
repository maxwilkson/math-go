package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {

	resultado := Soma(15, 15)

	assert.EqualValues(t, resultado, 30)
}
