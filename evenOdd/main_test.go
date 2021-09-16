package main

import "testing"

func TestIsEven(t *testing.T) {

	if !isEven(0) {
		t.Errorf("0 should be even")
	}

	if isEven(1) {
		t.Errorf("1 should be odd")
	}

	if !isEven(2) {
		t.Errorf("2 should be even")
	}

}
