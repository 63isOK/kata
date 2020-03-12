package door

import (
	"testing"
)

const want = `
1 0 0 1 0 0 0 0 1 0 
0 0 0 0 0 1 0 0 0 0 
0 0 0 0 1 0 0 0 0 0 
0 0 0 0 0 1 0 0 0 0 
0 0 0 0 0 0 0 0 1 0 
0 0 0 0 0 0 0 0 0 0 
0 0 0 1 0 0 0 0 0 0 
0 0 0 0 0 0 0 0 0 0 
1 0 0 0 0 0 0 0 0 0 
0 0 0 0 0 0 0 0 0 1 `

var tests = map[string]func() []bool{
	"doorFor":    doorFor,
	"doorFactor": doorFactor,
}

func TestAll(t *testing.T) {
	for name, f := range tests {
		ret := f()
		got := printString(t, ret)

		if want != got {
			t.Fatalf("\n%s, want:%sgot:%s\n", name, want, got)
		}
	}
}

func printString(t *testing.T, b []bool) string {
	t.Helper()

	ret := ""
	for i, x := range b {
		if i%10 == 0 {
			ret += "\n"
		}
		if x {
			ret += "1 "
		} else {
			ret += "0 "
		}
	}

	return ret
}

func BenchmarkDoorFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doorFor()
	}
}

func BenchmarkDoorFactor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doorFactor()
	}
}
