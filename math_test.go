package main

import "testing"

func TestSoma(t *testing.T) {
	if Soma(10, 10) != 20 {
		t.Error("Soma de 10 + 10 deve ser 20")
	}
	if Soma(15, 15) != 30 {
		t.Error("Soma de 15 + 15 deve ser 30")
	}
}
