package puzzle

import (
	"errors"
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

// Puzzle si a 4*4 puzzle
type Puzzle struct {
	data    [16]int
	x, y    int
	handler interface{} // Notify func
	// review  []int
	signal chan Signal
}

// New create a Puzzle
func New() *Puzzle {
	p := &Puzzle{data: [16]int{}}
	p.randomFill()

	go waitMoveEvent()

	return p
}

// ReStart is restart a new game
func (p *Puzzle) ReStart() (x, y int) {
	p.randomFill()

	// todo review
}

func (p *Puzzle) randomFill() {
}

func (p *Puzzle) waitMoveEvent() {

}

// Position get piece's position
func (p *Puzzle) Position() (x, y int) {
	return p.x, p.y
}

// Data get all info of puzzle
func (p *Puzzle) Data() [16]int {
	return p.data
}

// Notify is a event fire on move event done
// there are two way for notify:
//    func(int,int), app use x,y
//    func([16]int), app use whole puzzle data
func (p *Puzzle) Notify(handler interface{}) error {
	if !handler.(func(int, int)) && !handler.(func([16]int)) {
		return errors.New("callback format is incorrect")
	}

	p.handler = handler
}

// Send is send signal to Puzzle object
func (p *Puzzle) Send(s Signal) {

}

var std = New()

// Restart is restart game with default puzzle
func Restart() (x, y int) {
	std.ReStart()
}

// Position get piece's position of default puzzle
func Position() (x, y int) {
	return std.Position(x, y)
}

// Data get all info of default puzzle
func Data() [16]int {
	return std.Data()
}

// Notify is a event fire on move event done
// there are two way for notify:
//    func(int,int), app use x,y
//    func([16]int), app use whole puzzle data
func Notify(handler interface{}) error {
	return std.Notify(handler)
}

// Send is send a signal to default puzzle
func Send(s signal) {
	std.Send(s)
}
