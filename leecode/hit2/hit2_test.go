package hit2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHit2(t *testing.T) {
	input_arr := []int{3, 4, 2}
	input_target := 5
	want := []int{0, 2}
	got := Hit2(input_arr, input_target)
	assert.Equal(t, got, want)
}
