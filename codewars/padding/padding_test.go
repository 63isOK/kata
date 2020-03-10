package padding

import (
	"reflect"
	"testing"
)

func TestPadding(t *testing.T) {

	samples := []struct {
		i, w int
		out  []byte
	}{
		{2019, 3, []byte{'2', '0', '1', '9'}},
		{2019, 5, []byte{'0', '2', '0', '1', '9'}},
		{2019, -1, []byte{'2', '0', '1', '9'}},
	}

	for _, x := range samples {
		got := padding(x.i, x.w)
		want := x.out
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("want:%v, got: %v", want, got)
		}
	}
}
