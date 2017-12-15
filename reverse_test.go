package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct{
		in, want string
	}{
		{"!dlrow ,olleH", "Hello, world!"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _,aCase := range cases {
		got := Reverse(aCase.in)
		if got != aCase.want {
			t.Errorf("Reverse(%q) == %q, expects %q", aCase.in, got, aCase.want)
		}
	}
}
