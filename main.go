package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '⚾'

		sleep = time.Second / 20
	)

	var (
		px, py int
		vx, vy = 1, 1
		cell   rune
	)

	buf := make([]rune, width*height)

	board := make([][]bool, height)

	for i := range board {
		board[i] = make([]bool, width)
	}

	screen.Clear()

	for i := 0; i < 1200; i++ {
		buf = buf[:0]

		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		fmt.Println(px, py)

		for y := range board {
			for x := range board[0] {
				board[y][x] = false
			}
		}

		board[py][px] = true

		for y := range board {
			for x := range board[0] {
				if board[y][x] {
					cell = cellBall
				} else {
					cell = cellEmpty
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}
		screen.MoveTopLeft()
		time.Sleep(sleep)
		fmt.Print(string(buf))
	}

}
