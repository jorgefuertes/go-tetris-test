package main

import (
	"fmt"
	"os"
	"tetris/internal/board"
	"tetris/internal/piece"
	"tetris/internal/shared"

	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
)

func main() {
	p := piece.New()
	b := board.New(shared.Point{X: 16, Y: 24}, p.GetShape())

	redraw(b)
	for {
		char, k, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		switch char {
		case 'x', 'X':
			os.Exit(0)
		case 'q', 'Q':
			p.Left()
			if err := b.SetPieceShape(p.GetShape()); err != nil {
				p.Undo()
			} else {
				redraw(b)
			}
			continue
		case 'w', 'W':
			p.Right()
			if err := b.SetPieceShape(p.GetShape()); err != nil {
				p.Undo()
			} else {
				redraw(b)
			}
			continue
		}

		switch k {
		case keyboard.KeyArrowUp:
			b.PieceUp()
			redraw(b)
			continue
		case keyboard.KeyArrowDown:
			b.PieceDown()
			redraw(b)
			continue
		case keyboard.KeyArrowLeft:
			b.PieceLeft()
			redraw(b)
			continue
		case keyboard.KeyArrowRight:
			b.PieceRight()
			redraw(b)
			continue
		}
	}
}

func redraw(b *board.Board) {
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Println(b.Get())
	fmt.Print("Cursor: move, Q/W: rotate, X: exit: ")
}
