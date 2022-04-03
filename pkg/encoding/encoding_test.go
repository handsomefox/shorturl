package encoding

import (
	"testing"
)

func TestEncoding(t *testing.T) {
	encoder := NewEncoder()
	url := "https://twitch.tv"

	encoded := encoder.Encode(url)

	if encoded != "aHR0cHM6Ly90d2l0Y2gudHY" {
		t.Fatal("Encoding failed.")
	}
}

func TestDecoding(t *testing.T) {
	decoder := NewDecoder()
	b64 := "aHR0cHM6Ly90d2l0Y2gudHY"

	decoded, err := decoder.Decode(b64)

	if err != nil {
		t.Fatal(err)
	}

	if decoded != "https://twitch.tv" {
		t.Fatal("Decoding failed.")
	}
}
