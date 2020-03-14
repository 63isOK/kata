package puzzle

import (
	"math/rand"
	"time"
)

// move event
const (
	UP Signal = iota + 1
	DOWN
	LEFT
	RIGHT
	ReStart
	Quit
)

// Signal is move direction
type Signal int

// Pos is x,y of puzzle
type Pos struct {
	x, y int
}

// Puzzle si a 4*4 puzzle
type Puzzle struct {
	data           [16]int
	x, y           int
	notifyPosition chan<- Pos     // Notify channel
	notifyData     chan<- [16]int // Notify channel
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

// ReStart is restart a new game
func (p *Puzzle) ReStart() (x, y int) {
	p.init()

	// todo review

	return p.x, p.y
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

	//  1  2  3  4
	//  5  6  7  8
	//  9 10 11 12
	// 13 14 15 16
	// s = (y-1)*4 + x
	for i := 0; i < 16; i++ {
		if p.data[i] == 0 {
			i++
			if i < 5 {
				p.y = 1
				p.x = i
			} else if i < 9 {
				p.y = 2
				p.x = i - 4
			} else if i < 13 {
				p.y = 3
				p.x = i - 8
			} else {
				p.y = 4
				p.x = i - 12
			}
			break
		}
	}

	p.isDone = false
}

func (p *Puzzle) waitMoveEvent() {
	for {
		select {
		case s := <-p.signal:
			if s == ReStart {
				p.ReStart()
				continue
			}

			if p.isDone {
				continue
			}

			if (s == UP && p.y == 1) ||
				(s == DOWN && p.y == 4) ||
				(s == LEFT && p.x == 1) ||
				(s == RIGHT && p.x == 4) {
				continue
			}

			switch s {
			case UP:
				p.data[p.y*4+p.x-5], p.data[p.y*4+p.x-9] = p.data[p.y*4+p.x-9], p.data[p.y*4+p.x-5]
				p.y--
			case DOWN:
				p.data[p.y*4+p.x-5], p.data[p.y*4+p.x-1] = p.data[p.y*4+p.x-1], p.data[p.y*4+p.x-5]
				p.y++
			case LEFT:
				p.data[p.y*4+p.x-5], p.data[p.y*4+p.x-6] = p.data[p.y*4+p.x-6], p.data[p.y*4+p.x-5]
				p.x--
			case RIGHT:
				p.data[p.y*4+p.x-5], p.data[p.y*4+p.x-4] = p.data[p.y*4+p.x-4], p.data[p.y*4+p.x-5]
				p.x++
			}

			if p.notifyPosition != nil {
				p.notifyPosition <- Pos{p.x, p.y}
			}
			if p.notifyData != nil {
				p.notifyData <- p.data
			}

			if p.x == 4 && p.y == 4 {
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

	p.done <- 1
}

// Position get piece's position
func (p *Puzzle) Position() (x, y int) {
	return p.x, p.y
}

// Data get all info of puzzle
func (p *Puzzle) Data() [16]int {
	return p.data
}

// NotifyPosition is a event fire on move event done
func (p *Puzzle) NotifyPosition(c chan<- Pos) {
	p.notifyPosition = c
}

// NotifyData is a event fire on move event done
func (p *Puzzle) NotifyData(c chan<- [16]int) {
	p.notifyData = c
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

// Restart is restart game with default puzzle
func Restart() (x, y int) {
	return std.ReStart()
}

// Position get piece's position of default puzzle
func Position() (x, y int) {
	return std.Position()
}

// Data get all info of default puzzle
func Data() [16]int {
	return std.Data()
}

// NotifyPosition is a event fire on move event done
func NotifyPosition(c chan<- Pos) {
	std.NotifyPosition(c)
}

// NotifyData is a event fire on move event done
func NotifyData(c chan<- [16]int) {
	std.NotifyData(c)
}

// Send is send a signal to default puzzle
func Send(s Signal) {
	std.Send(s)
}
