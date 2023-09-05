package piece

import (
	"tetris/internal/shared"
)

type Piece struct {
	orientation uint8
	was         uint8
}

func New() *Piece {
	return &Piece{}
}

func (p *Piece) Left() {
	p.was = p.orientation
	if p.orientation == 0 {
		p.orientation = 3
		return
	}
	p.orientation--
}

func (p *Piece) Right() {
	p.was = p.orientation
	p.orientation++
	if p.orientation > 3 {
		p.orientation = 0
	}
}

func (p *Piece) Undo() {
	current := p.orientation
	p.orientation = p.was
	p.was = current
}

func (p Piece) GetShape() shared.Mask {
	switch p.orientation {
	case 0:
		// 🔳
		// 🔳🔳🔳
		return shared.Mask{
			{true},
			{true, true, true},
		}
	case 1:
		// 🔳🔳
		// 🔳
		// 🔳
		return shared.Mask{
			{true, true},
			{true},
			{true},
		}
	case 2:
		// 🔳🔳🔳
		//     🔳
		return shared.Mask{
			{true, true, true},
			{false, false, true},
		}
	default:
		//   🔳
		//   🔳
		// 🔳🔳
		return shared.Mask{
			{false, true},
			{false, true},
			{true, true},
		}
	}
}
