package main

import (
	"math/rand"

	termbox "github.com/nsf/termbox-go"

	//	"strconv"
	"time"
)

// WIDTH is width of fire
const WIDTH = 16

// HEIGHT is height of file
const HEIGHT = 16

// FLUCTUATION is fluctuation of fire
const FLUCTUATION = 5

func main() {
	var m [WIDTH][HEIGHT]int
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			m[x][y] = 0
		}
	}

	colors := ([]termbox.Attribute{
		0,
		197,
		203,
		209,
		215,
		221,
		227,
		228,
		229,
		230,
		231})

	rand.Seed(time.Now().UnixNano())

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	ch := make(chan termbox.Event)
	go keyEvent(ch)

loop:
	for {

		select {
		case ev := <-ch:
			switch ev.Type {
			case termbox.EventKey:
				if ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyEsc {
					break loop
				}
			}
			break
		default:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			fire(&m, colors)
			drawFire(m, colors)
			termbox.Flush()
		}
	}
}

func fire(m *[WIDTH][HEIGHT]int, colors []termbox.Attribute) {
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {

			if y < HEIGHT-2 {
				if x == WIDTH-1 {
					m[x][y] = (m[x][y] + m[x-1][y+1] + m[x][y+1] + m[x][y+2]) / 5
				} else if x == 0 {
					m[x][y] = (m[x][y] + m[x][y+1] + m[x+1][y+1] + m[x][y+2]) / 5
				} else {
					m[x][y] = (m[x][y] + m[x-1][y+1] + m[x][y+1] + m[x+1][y+1] + m[x][y+2]) / 5
				}
			} else if y < HEIGHT-1 {
				if x == WIDTH-1 {
					m[x][y] = (m[x][y] + m[x-1][y+1] + m[x][y+1]) / 5
				} else if x == 0 {
					m[x][y] = (m[x][y] + m[x][y+1] + m[x+1][y+1]) / 5
				} else if x < 3 || x > WIDTH-4 {
					m[x][y] = (m[x][y] + m[x-1][y+1] + m[x][y+1] + m[x+1][y+1]) / 5
				}
			} else {
				if x < 3 || x > WIDTH-4 {
					m[x][y] = m[x][y] / 5
				}
			}
			if rand.Intn(FLUCTUATION) == 1 {
				m[x][y] = m[x][y] - 1
			}
			if m[x][y] < 0 {
				m[x][y] = 0
			}
		}
	}

	for x := 5; x < WIDTH-5; x++ {
		m[x][HEIGHT-1] = rand.Intn(3) + len(colors) - 3

	}
	for x := 4; x < WIDTH-4; x++ {
		m[x][HEIGHT-2] = rand.Intn(3) + len(colors) - 3
	}

}

func keyEvent(ch chan termbox.Event) {
	for {
		ch <- termbox.PollEvent()
	}
}

func drawFire(m [WIDTH][HEIGHT]int, colors []termbox.Attribute) {
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			termbox.SetCell(x+2, y+2, ' ', termbox.ColorDefault, colors[m[x][y]])
		}
	}
}
