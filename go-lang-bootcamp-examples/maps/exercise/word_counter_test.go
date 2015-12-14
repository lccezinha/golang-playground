package main 

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
    expected := map[string]int{
        "xunda": 5,
        "maneiro": 7,
    }

    text := "xunda maneiro"
    result := WordCount(text)

    assert.Equal(t, result, expected, "Deve retornar o contador corretamente!")
}