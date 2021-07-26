package main

import (
	"testing"
)

func TestHelloServerLogic(t *testing.T) {
	serverResponse := HelloServerLogic("Tom")
	if "Hello, Tom!" != serverResponse {
		t.Errorf("HelloServerLogic unexpected response %s", serverResponse)
	}
}
