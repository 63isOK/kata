package main

import (
	"errors"
	"math/rand"
	"time"
)

// Signal is app signal which is used to control the puzzle
type Signal int

// app signal
const (
	UP Signal = iota + 1
	DOWN
	LEFT
	RIGHT
	NEW
	REVIEW
	Quit
)

// Data is puzzle's data
type Data [16]int

// Puzzle si a 4*4 puzzle
type Puzzle struct {
	data       Data
	notify     chan<- Data
	reviewData []Data
	reviewReay bool
	signal     chan Signal
	isDone     bool
	done       chan interface{}
}

// New create a Puzzle
func New(notify chan<- Data) *Puzzle {
	if notify == nil {
		return nil
	}

	p := &Puzzle{
		data:   Data{},
		notify: notify,
		signal: make(chan Signal),
		done:   make(chan interface{})}
	p.init()

	go p.waitMoveEvent()

	p.Send(NEW)

	return p
}

// reStart is restart a new game
func (p *Puzzle) reStart() {
	p.init()

	p.notify <- p.data
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
	p.reviewReay = false
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
				p.reStart()
				continue
			}

			if p.isDone {
				if s == REVIEW {
					for _, x := range p.reviewData {
						p.notify <- x
						time.Sleep(time.Second)
					}
				}
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

				p.notify <- p.data
				continue
			}

			if !p.reviewReay {
				p.reviewData = make([]Data, 0)
				p.reviewReay = true
				p.reviewData = append(p.reviewData, p.data)
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

			p.notify <- p.data
			p.reviewData = append(p.reviewData, p.data)

			if p.data[15] == 0 {
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

// Send is send signal to Puzzle object
func (p *Puzzle) Send(s Signal) {
	p.signal <- s
}

// Done get done channel
func (p *Puzzle) Done() <-chan interface{} {
	return p.done
}
