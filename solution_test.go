package solutions

import "testing"

func TestGetMessage(t *testing.T) {
	expected := string([]rune{72, 101, 108, 108, 111, 32, 128506, 65039, 32, 33})

	got := GetMessage()

	if got != expected {
		t.Errorf("got: %v not equal expected: %v\n", got, expected)
	}
}
