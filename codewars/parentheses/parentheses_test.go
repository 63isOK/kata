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

type targetFunc func(string) bool

var targets = []targetFunc{valid}

func TestAll(t *testing.T) {
	for _, data := range tests {
		for _, f := range targets {
			got := f(data.in)
			if data.out != got {
				t.Fatalf("%q, want:%t, got:%t", data.in, data.out, got)
			}
		}
	}
}

func BenchmarkValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valid("(())((()())())")
	}
}
