package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if HelloWorld() != "Hello World" {
		t.Fatalf("HelloWorld func is not return correctly")
	}
}
