package main

import (
	"fmt"
)

func main() {
	dataCh := make(chan Data)
	p := New(dataCh)
	doneCh := p.Done()

	go func(p *Puzzle) {
		for {
			var c byte
			if _, err := fmt.Scan(&c); err == nil {
				switch c {
				case 'h':
					p.Send(LEFT)
				case 'j':
					p.Send(DOWN)
				case 'k':
					p.Send(UP)
				case 'l':
					p.Send(RIGHT)
				case 'r':
					p.Send(REVIEW)
				case 'a':
					p.Send(NEW)
				}
			}
		}
	}(p)

	for {
		select {
		case data := <-dataCh:
			fmt.Print("\f")
			fmt.Println("[hjkl] move:Left/Down/Up/Right")
			fmt.Println("[r] review")
			fmt.Println("[a] play again")
			fmt.Println()

			for i := 0; i < 16; i++ {
				fmt.Print(data[i])
				fmt.Print(" ")
				if i%4 == 3 {
					fmt.Println()
				}
			}
		case <-doneCh:
			fmt.Print("\a")
		}
	}
}
