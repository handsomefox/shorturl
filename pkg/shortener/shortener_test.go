package shortener

import (
	"testing"
)

func TestHash(t *testing.T) {
	hex := hash("string")

	if hex != "6531c93affc" {
		t.Fatal()
	}
}

func TestMake(t *testing.T) {
	hex, link := Make("string")

	if hex != "6531c93affc" {
		t.Fatal()
	}

	if link != "localhost:3000/u/6531c93affc" {
		t.Fatal()
	}
}
