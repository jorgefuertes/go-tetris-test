package board

import (
	"sync"
	"tetris/internal/shared"
)

type Direction string
const Left Direction = "left"
const Right Direction = "right"
const Down Direction = "down"

type Board struct {
	size     shared.Point
	data     shared.Mask
	position shared.Point
	shape    shared.Mask
	mux      *sync.Mutex
}

func New(size shared.Point, shape shared.Mask) *Board {
	return &Board{
		size:     size,
		data:     newData(size),
		position: shared.Point{X: 0, Y: 0},
		shape:    shape,
		mux:      new(sync.Mutex),
	}
}

func (b *Board) SetPieceShape(shape shared.Mask) error {
	b.mux.Lock()
	defer b.mux.Unlock()
	if !b.IsSpaceFor(shape) {
		return ErrNotEnoughSpace
	}
	b.shape = shape
	return nil
}

func (b Board) IsSpaceFor(shape shared.Mask) bool {
	if shape.Height()+b.position.Y > b.size.Y {
		return false
	}
	if shape.Width()+b.position.X > b.size.X {
		return false
	}
	return true
}
