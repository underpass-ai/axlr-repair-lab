package lab

import "testing"

func TestWordCount(t *testing.T) {
	cases := map[string]int{"hola mundo": 2, "hola  mundo": 2, "  hola mundo  ": 2, "": 0, "uno": 1}
	for text, want := range cases {
		if got := WordCount(text); got != want {
			t.Errorf("WordCount(%q) = %d, want %d", text, got, want)
		}
	}
}
