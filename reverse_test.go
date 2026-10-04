package lab

import "testing"

func TestReverse(t *testing.T) {
	cases := map[string]string{"hola": "aloh", "": "", "año": "oña", "日本": "本日"}
	for text, want := range cases {
		if got := Reverse(text); got != want {
			t.Errorf("Reverse(%q) = %q, want %q", text, got, want)
		}
	}
}
