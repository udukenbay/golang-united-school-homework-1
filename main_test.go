package main

import "testing"

func TestGetMessage(t *testing.T) {
	s := GetMessage()
	if len(s) == 0 {
		t.Error("Error")
	}
}
