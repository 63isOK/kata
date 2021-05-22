package hit2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHit2(t *testing.T) {
	hit2(t, Hit2)
}

func TestHit2Hash(t *testing.T) {
	hit2(t, Hit2HashMap)
}

func hit2(t *testing.T, impl func([]int, int) []int) {
	input_arr := []int{3, 4, 2}
	input_target := 5
	want := []int{0, 2}
	got := impl(input_arr, input_target)
	assert.Equal(t, got, want)
}
