package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		given, wanted string
	}{
		{"Hello world", "dlrow olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, caseI := range cases {
		actual := Reverse(caseI.given)
		if actual != caseI.wanted {
			t.Error("Reverse(%q) == %q, wanted %q", caseI.given, actual, caseI.wanted)
		}
	}
}
