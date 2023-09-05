package board

func (b *Board) setPositionX(x int) {
	if x < 0 || x > b.size.X-b.shape.Width() {
		return
	}
	b.mux.Lock()
	defer b.mux.Unlock()
	b.position.X = x
}

func (b *Board) setPositionY(y int) {
	if y < 0 || y > b.size.Y-b.shape.Height() {
		return
	}
	b.mux.Lock()
	defer b.mux.Unlock()
	b.position.Y = y
}

func (b *Board) PieceDown() {
	b.setPositionY(b.position.Y + 1)
}

func (b *Board) PieceUp() {
	b.setPositionY(b.position.Y - 1)
}
func (b *Board) PieceRight() {
	b.setPositionX(b.position.X + 1)
}

func (b *Board) PieceLeft() {
	b.setPositionX(b.position.X - 1)
}
