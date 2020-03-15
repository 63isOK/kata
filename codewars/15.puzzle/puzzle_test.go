package puzzle

import (
	"testing"
	"time"
)

func TestConStruct(t *testing.T) {
	allData := make(chan Data)
	p := New(allData)
	go func() {
		for i := 0; i < 100; i++ {

			p.reStart()
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

func TestDone(t *testing.T) {
	dataCh := make(chan Data)
	p := New(dataCh)
	doneCh := p.Done()
	go func(t *testing.T) {
		for {
			select {
			case <-doneCh:
				p.Send(REVIEW)
				return
			case <-time.After(time.Millisecond * 100):
				t.Fatal("it shuoud be done,but not")
			}
		}
	}(t)

	p.data = Data{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 0, 15}
	p.Send(RIGHT)
	<-dataCh

	time.Sleep(time.Millisecond * 200)

	data := <-dataCh
	dataReview := Data{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 0, 15}
	compareData(t, dataReview, data)
	data = <-dataCh
	dataReview = Data{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0}
	compareData(t, dataReview, data)
}

func compareData(t *testing.T, a, b Data) {
	t.Helper()

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Fatalf("review,want:%v,got:%v", a, b)
		}
	}
}
