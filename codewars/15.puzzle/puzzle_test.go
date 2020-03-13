package puzzle

import (
	"testing"
)

func TestConStruct(t *testing.T) {
	p := New()
	checkDuplicate(t, p)
	p.ReStart()
	checkDuplicate(t, p)
}

func checkDuplicate(t *testing.T, p *Puzzle) {
	t.Helper()

	data = p.GetData()

	checkCount := 0

	for i := 0; i < 16; i++ {
		for _, x := range data {
			if i == x {
				checkCount++
				continue
			}
		}
	}

	if checkCount != 16 {
		t.Fatalf("there is a error in New():%v", data)
	}
}
