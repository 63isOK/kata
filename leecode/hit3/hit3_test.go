package hit3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHit3(t *testing.T) {
	input_arr := []int{3, 4, 2, -1, -1, 0, -2, 1, 0, 0, 0}
	want := [][]int{
		[]int{-2, -1, 3},
		[]int{-2, 0, 2},
		[]int{-1, -1, 2},
		[]int{-1, 0, 1},
		[]int{0, 0, 0}}
	got := hit3(input_arr)
	assert.ElementsMatch(t, got, want)
}

func TestHit32(t *testing.T) {
	input_arr := []int{-1, 0, 1, 2, -1, -4}
	want := [][]int{
		[]int{-1, -1, 2},
		[]int{-1, 0, 1}}
	got := hit32(input_arr)
	assert.ElementsMatch(t, got, want)
}

func TestHit3Web(t *testing.T) {
	input_arr := []int{-1, 0, 1, 0}
	want := [][]int{[]int{-1, 0, 1}}
	got := hit3(input_arr)
	assert.ElementsMatch(t, got, want)
}

func BenchmarkHit3(b *testing.B) {
	input_arr := []int{3, 4, 2, -1, -1, 0, -2, 1, 0, 0, 0}
	for i := 0; i < b.N; i++ {
		hit3(input_arr)
	}
}

func TestHit3Point(t *testing.T) {
	input_arr := []int{-1, 0, 1, 2, -1, -4}
	want := [][]int{
		[]int{-1, -1, 2},
		[]int{-1, 0, 1}}
	got := hit3Point(input_arr)
	assert.ElementsMatch(t, got, want)
}