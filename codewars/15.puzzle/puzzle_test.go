package puzzle

import (
	"testing"
)

func TestConStruct(t *testing.T) {
	allData := make(chan Data)
	p := New(allData)
	go func() {
		for i := 0; i < 100; i++ {

			p.ReStart()
		}
		close(allData)
	}()

	for {
		select {
		case data, ok := <-allData:
			if !ok {
				return
			}

			checkDuplicate(t, data)
		}
	}
}

func checkDuplicate(t *testing.T, data Data) {
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
	allData := make(chan Data)
	p := New(allData)

	p.Send(UP)
	<-allData
	p.Send(UP)
	<-allData
	p.Send(UP)
	<-allData
	p.Send(LEFT)
	<-allData
	p.Send(LEFT)
	<-allData
	p.Send(LEFT)
	data := <-allData
	checkSwap(t, data, "init status", 0)

	p.Send(LEFT)
	data = <-allData
	checkSwap(t, data, "unreach left", 0)
	p.Send(UP)
	data = <-allData
	checkSwap(t, data, "unreach up", 0)

	p.Send(RIGHT)
	data = <-allData
	checkSwap(t, data, "1 right", 1)
	p.Send(DOWN)
	data = <-allData
	checkSwap(t, data, "1 right, 1 down", 5)
	p.Send(RIGHT)
	data = <-allData
	checkSwap(t, data, "2 right, 1 down", 6)
	p.Send(DOWN)
	data = <-allData
	checkSwap(t, data, "2 right, 2 down", 10)
	p.Send(RIGHT)
	data = <-allData
	checkSwap(t, data, "3 right, 2 down", 11)
	p.Send(DOWN)
	data = <-allData
	checkSwap(t, data, "3 right, 3 down", 15)

	p.Send(RIGHT)
	data = <-allData
	checkSwap(t, data, "unreach right", 15)
	p.Send(DOWN)
	data = <-allData
	checkSwap(t, data, "unreadch down", 15)
}

func checkSwap(t *testing.T, data Data, do string, index int) {
	t.Helper()

	if data[index] != 0 {
		t.Fatalf("%s,want:0,data[%d] get:(%v)", do, index, data)
	}
}
