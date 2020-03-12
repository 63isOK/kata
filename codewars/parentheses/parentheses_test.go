package parentheses

import (
	"testing"
)

type tester struct {
	in  string
	out bool
}

var tests = []tester{
	{"()", true},
	{")(()))", false},
	{"(", false},
	{"(())((()())())", true},
}

func TestAll(t *testing.T) {
	for _, x := range tests {
		got := valid(x.in)
		if x.out != got {
			t.Fatalf("%q, want:%t, got:%t", x.in, x.out, got)
		}
	}
}

func BenchmarkValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valid("(())((()())())")
	}
}
