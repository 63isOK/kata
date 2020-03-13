package puzzle

import (
	"testing"
)

var px, py = 0, 0
var pdata = 0

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

func TestStd(t *testing.T) {
	Send(UP)
	Send(UP)
	Send(UP)
	Send(LEFT)
	Send(LEFT)
	Send(LEFT)

	Notify(func(x, y int) {
		px, py = x, y
	})

	checkSwap(t, "init status", 1, 1, 0)
	Send(LEFT)
	checkSwap(t, "unreach left", 1, 1, 0)
	Send(UP)
	checkSwap(t, "unreach up", 1, 1, 0)

	Send(RIGHT)
	checkSwap(t, "1 right", 2, 1, 1)
	Send(DOWN)
	checkSwap(t, "1 right, 1 down", 2, 2, 5)
	Send(RIGHT)
	checkSwap(t, "2 right, 1 down", 3, 2, 6)
	Send(DOWN)
	checkSwap(t, "2 right, 2 down", 3, 3, 10)

	Notify(func(data [16]int) {
		pdata = data
	})

	Send(RIGHT)
	checkSwap(t, "3 right, 2 down", 4, 3, 14)
	Send(DOWN)
	checkSwap(t, "3 right, 3 down", 4, 4, 15)

	Send(RIGHT)
	checkSwap(t, "unreach right", 4, 4, 15)
	Send(DOWN)
	checkSwap(t, "unreadch down", 4, 4, 15)
}

func checkSwap(t *testing.T, do string, x, y, index int) {
	t.Helper()

	posx, posy := Position()
	if posx != x || posy != y {
		t.Fatalf("%s,(x,y) must be %d,%d", do, x, y)
	}
	data := Data()
	if data[index] != 0 {
		t.Fatalf("%s,data[index] must be 0", do, index)
	}

	if std.handler.(func(int, int)) {
		if px != x || py != y {
			t.Fatalf("%s,want:(%d,%d),callback get:(%d,%d)", do, x, y, px, py)
		}
	} else if std.handler.(func([16]int)) {
		if pdata[index] != 0 {
			t.Fatalf("%s,want:0,pdata[%d] get:(%d)", do, index, pdata[index])
		}
	}
}
