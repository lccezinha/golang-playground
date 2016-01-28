package main

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestEmptyStack(t *testing.T) {
  stack := Stack{}
  result := stack.Empty()

  assert.Equal(t, true, result, "must be true")
}

func TestPush(t *testing.T) {
  stack := Stack{}

  stack.Push(1)
  stack.Push("Xunda")

  empty := stack.Empty()
  size := stack.Size()

  assert.Equal(t, empty, false, "can not be empty")
  assert.Equal(t, size, 2, "size be 2")
}

func TestRemoveItemSuccess(t *testing.T) {
  assert := assert.New(t)
  stack := Stack{}

  stack.Push(1)
  stack.Push("xunda")
  result, err := stack.Pop()

  empty := stack.Empty()
  size := stack.Size()

  assert.Equal(empty, false, "Have item")
  assert.Equal(1, size)
  assert.Equal(result, "xunda")
  assert.Nil(err)
}

func TestRemoveItemError(t *testing.T) {
  assert := assert.New(t)
  stack := Stack{}

  result, err := stack.Pop()

  assert.Nil(result)
  if assert.NotNil(err) {
    assert.Equal("Empty Stack", err.Error())
  }
}

func TestSize(t *testing.T) {
  stack := Stack{}

  stack.Push(1)

  result := stack.Size()

  assert.Equal(t, 1, result, "must be 1")
}