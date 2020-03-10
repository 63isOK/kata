package padding

import (
	"testing"
)

func TestPadding(t *testing.T) {

	samples := []struct {
		i, w int
		out  string
		fun  func(int, int) string
	}{
		{2018, 3, "2018", padding},
		{2017, 5, "02017", padding},
		{2016, -1, "2016", padding},
		{2015, 3, "2015", paddingNoStd},
		{2014, 5, "02014", paddingNoStd},
		{2013, -1, "2013", paddingNoStd},
		{2012, 3, "2012", paddingStd},
		{2011, 5, "02011", paddingStd},
		{2010, -1, "2010", paddingStd},
		{2009, 3, "2009", paddingNoStd2},
		{2008, 5, "02008", paddingNoStd2},
		{2007, -1, "2007", paddingNoStd2},
	}

	for _, x := range samples {
		got := x.fun(x.i, x.w)
		want := x.out
		if want != got {
			t.Fatalf("want:%v, got: %v", want, got)
		}
	}
}

var samples = []struct {
	I, W int
}{
	{2018, 3},
	{2017, 5},
	{2016, -1},
}

func BenchmarkPadding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range samples {
			padding(x.I, x.W)
		}
	}
}

func BenchmarkPaddingNoStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range samples {
			paddingNoStd(x.I, x.W)
		}
	}
}

func BenchmarkPaddingNoStd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range samples {
			paddingNoStd2(x.I, x.W)
		}
	}
}

func BenchmarkPaddingStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, x := range samples {
			paddingStd(x.I, x.W)
		}
	}
}
