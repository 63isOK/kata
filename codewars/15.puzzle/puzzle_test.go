package puzzle

import (
	"testing"
)

var allData chan Data

func init() {
	allData = make(chan Data)
}

func TestConStruct(t *testing.T) {
	p := New()
	data := p.Start()
	checkDuplicate(t, p, data)
	data = p.ReStart()
	checkDuplicate(t, p, data)
}

func checkDuplicate(t *testing.T, p *Puzzle, data Data) {
	t.Helper()

	checkCount := 0

	for i := 0; i < 16; i++ {
		for _, x := range data {
			if i == x {
				checkCount++
				break
			}
		}
	}

	if checkCount != 16 {
		t.Fatalf("there is a duplicate error in puzzle :%v", data)
	}
}

func TestStd(t *testing.T) {
	Send(UP)
	Send(UP)
	Send(UP)
	Send(LEFT)
	Send(LEFT)
	Send(LEFT)
	checkSwap(t, "init status", 0)

	Notify(allData)
	Send(LEFT)
	checkSwap(t, "unreach left", 0)
	Send(UP)
	checkSwap(t, "unreach up", 0)

	Send(RIGHT)
	checkSwap(t, "1 right", 1)
	Send(DOWN)
	checkSwap(t, "1 right, 1 down", 5)
	Send(RIGHT)
	checkSwap(t, "2 right, 1 down", 6)
	Send(DOWN)
	checkSwap(t, "2 right, 2 down", 10)
	Send(RIGHT)
	checkSwap(t, "3 right, 2 down", 11)
	Send(DOWN)
	checkSwap(t, "3 right, 3 down", 15)

	Send(RIGHT)
	checkSwap(t, "unreach right", 15)
	Send(DOWN)
	checkSwap(t, "unreadch down", 15)
}

func checkSwap(t *testing.T, do string, index int) {
	t.Helper()

	if std.notify != nil {
		pdata := <-allData
		if pdata[index] != 0 {
			t.Fatalf("%s,want:0,pdata[%d] get:(%v)", do, index, pdata)
		}
	}
}
