package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertPrint := func(t *testing.T, want, got string) {
		t.Helper()

		if want != got {
			t.Errorf("want: %s, got: %s", want, got)
		}
	}

	t.Run("nobody", func(t *testing.T) {
		want := "hello world"
		got := hello("")

		assertPrint(t, want, got)
	})

	t.Run("tom", func(t *testing.T) {
		want := "hello tom"
		got := hello("tom")

		assertPrint(t, want, got)
	})
}
