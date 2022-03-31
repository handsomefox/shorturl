package shorturl

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
	hex := Make("string")

	if hex != "6531c93affc" {
		t.Fatal()
	}
}
