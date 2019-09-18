package main

import "testing"

func Test_getMessage(t *testing.T) {
	msg := getMessage("sample")
	if msg != "Hello World from the cluster namespace 'sample'" {
		t.Errorf("Wrong message: %s", msg)
	}
}
