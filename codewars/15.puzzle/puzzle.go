package puzzle

import (
	"errors"
	"math/rand"
	"time"
)

// move event
const (
	UP Signal = iota + 1
	DOWN
	LEFT
	RIGHT
	NEW
	Quit
)

// Signal is move direction
type Signal int

// Data is puzzle's data
type Data [16]int

// Puzzle si a 4*4 puzzle
type Puzzle struct {
	data   Data
	notify chan<- Data
	// review  []int
	signal chan Signal
	isDone bool
	done   chan interface{}
}

// New create a Puzzle
func New() *Puzzle {
	p := &Puzzle{data: [16]int{}, signal: make(chan Signal), done: make(chan interface{})}
	p.init()

	go p.waitMoveEvent()

	return p
}

// Start is start a new game
func (p *Puzzle) Start() Data {
	return p.ReStart()
}

// ReStart is restart a new game
func (p *Puzzle) ReStart() Data {
	p.init()

	// todo review

	return p.data
}

func (p *Puzzle) init() {
	for i := 0; i < 16; i++ {
		p.data[i] = i
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 16; i > 0; i-- {
		random := r.Intn(i)
		p.data[i-1], p.data[random] = p.data[random], p.data[i-1]
	}

	p.isDone = false
}

func (p *Puzzle) getSwapIndex() (int, error) {
	for i := range p.data {
		if p.data[i] == 0 {
			return i, nil
		}
	}

	return 0, errors.New("invalid data")
}

func (p *Puzzle) waitMoveEvent() {
	for {
		select {
		case s := <-p.signal:
			if s == NEW {
				p.ReStart()
				continue
			}

			if p.isDone {
				continue
			}

			index, err := p.getSwapIndex()
			if err != nil {
				panic(err.Error())
			}

			if (s == UP && index < 4) ||
				(s == DOWN && index > 11) ||
				(s == LEFT && index%4 == 0) ||
				(s == RIGHT && index%4 == 3) {

				if p.notify != nil {
					p.notify <- p.data
				}
				continue
			}

			switch s {
			case UP:
				p.data[index], p.data[index-4] = p.data[index-4], p.data[index]
			case DOWN:
				p.data[index], p.data[index+4] = p.data[index+4], p.data[index]
			case LEFT:
				p.data[index], p.data[index-1] = p.data[index-1], p.data[index]
			case RIGHT:
				p.data[index], p.data[index+1] = p.data[index+1], p.data[index]
			}

			if p.notify != nil {
				p.notify <- p.data
			}

			if index == 15 {
				p.checkDone()
			}
		}
	}
}

func (p *Puzzle) checkDone() {
	for i, x := range p.data[:15] {
		if i+1 != x {
			return
		}
	}

	p.isDone = true
	p.done <- 1
}

// Notify is a event fire on move event done
func (p *Puzzle) Notify(c chan<- Data) {
	p.notify = c
}

// Send is send signal to Puzzle object
func (p *Puzzle) Send(s Signal) {
	p.signal <- s
}

// Done get done channel
func (p *Puzzle) Done() <-chan interface{} {
	return p.done
}

var std = New()

// Start is start a new game
func Start() Data {
	return std.Start()
}

// ReStart is restart a new game
func ReStart() Data {
	return std.ReStart()
}

// Notify is a event fire on move event done
func Notify(c chan<- Data) {
	std.Notify(c)
}

// Send is send signal to Puzzle object
func Send(s Signal) {
	std.Send(s)
}

// Done get done channel
func Done() <-chan interface{} {
	return std.Done()
}
