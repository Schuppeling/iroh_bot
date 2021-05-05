package quotes

import (
	"testing"
)

func TestRandomQuotes(t *testing.T) {
	quote := GetRandomQuote()

	if quote == "" {
		t.Fatalf("Expected quote but no quote was returned")
	}
}

func TestReadQuotes(t *testing.T) {
	readQuotes()

	if len(quoteMap) == 0 {
		t.Fatalf("Expected quote map to contain quotes but contained 0 quotes")
	}
}
