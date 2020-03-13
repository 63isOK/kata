package puzzle

// Puzzle si a 4*4 puzzle
type Puzzle struct {
	data    [16]int
	x, y    int
	handler interface{} // Notify func
	review  []int
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

// GetPiecePosition get piece's position
func (p *Puzzle) GetPiecePosition() (x, y int) {
	return 0, 0
}

// GetData get all info of puzzle
func (p *Puzzle) GetData() [16]int {
	return p.data
}

// Notify is a event fire on move event done
// there are two way for notify:
//    func(int,int), app use x,y
//    func([16]int), app use whole puzzle data
func (p *Puzzle) Notify(handler interface{}) {
	p.handler = handler
}

var std = New()

// Start is start game with default puzzle
func Start() {
	std.Start()
}

// Restart is restart game with default puzzle
func Restart() {
	std.ReStart()
}

// GetPiecePosition get piece's position of default puzzle
func GetPiecePosition() (x, y int) {
	return std.GetPiecePosition(x, y)
}

// GetData get all info of default puzzle
func GetData() [16]int {
	return std.GetData()
}

// Notify is a event fire on move event done
// there are two way for notify:
//    func(int,int), app use x,y
//    func([16]int), app use whole puzzle data
func Notify(handler interface{}) {
	std.Notify(handler)
}
