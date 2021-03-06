package rangeint

import "testing"

type tester struct {
	in  []int
	out string
}

var tests = []tester{
	{[]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}, "-6,-3-1,3-5,7-11,14,15,17-20"},
	{[]int{0, 1, 2, 4, 6, 7, 8, 11, 12, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 27, 28, 29, 30, 31, 32, 33, 35, 36, 37, 38, 39}, "0-2,4,6-8,11,12,14-25,27-33,35-39"},
	{[]int{78, 82, 83, 84, 87, 94, 97, 98, 99, 100, 101, 102, 104, 105}, "78,82-84,87,94,97-102,104,105"},
}

func TestAll(t *testing.T) {
	for _, x := range tests {
		got := solution(x.in)
		if got != x.out {
			t.Fatalf("want:%q, got:%q", x.out, got)
		}

		got2 := solution2(x.in)
		if got2 != x.out {
			t.Fatalf("want:%q, got:%q", x.out, got)
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20})
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution2([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20})
	}
}
