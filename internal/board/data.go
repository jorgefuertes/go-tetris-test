package board

import "tetris/internal/shared"

func newData(size shared.Point) shared.Mask {
	data := make(shared.Mask, size.Y)
	for i := range data {
		data[i] = make(shared.MaskRow, size.X)
	}
	return data
}
