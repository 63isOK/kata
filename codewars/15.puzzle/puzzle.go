package puzzle

// Puzzle si a 4*4 puzzle
type Puzzle struct {
	data   [16]int
	review []int
}

// move event
const (
	UP    = 8
	DOWN  = 2
	LEFT  = 4
	RIGHT = 6
)

var event chan int

// New create a Puzzle
func New() *Puzzle {
	p = &Puzzle{}
	return p.init()
}

func (p *Puzzle) init() *Puzzle {
	return p
}

// ReStart is restart a new game
func (p *Puzzle) ReStart() {
	Start()
}

// Start is game start
func (p *Puzzle) Start() {
}
