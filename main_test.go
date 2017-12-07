package main

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	actual := getMessage()
	excepted := "Hello, World!"
	if actual != excepted {
		t.Errorf("got %v\n want %v", actual, excepted)
	}
}
