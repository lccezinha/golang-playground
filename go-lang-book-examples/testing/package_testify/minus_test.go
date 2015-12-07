package main 

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMinus(t *testing.T) {
    n1 := 10
    n2 := 5
    expected := 5
    result := minus(n1, n2)

    assert.Equal(t, result, expected, "Devem ser iguais.")
}