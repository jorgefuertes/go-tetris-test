package board

import "fmt"

func (b *Board) clear() {
	b.mux.Lock()
	defer b.mux.Unlock()
	for row := 0; row < b.size.Y; row++ {
		for col := 0; col < b.size.X; col++ {
			b.data[row][col] = false
		}
	}
}

func (b *Board) paint() {
	b.mux.Lock()
	defer b.mux.Unlock()
	fmt.Printf("X: %d Y: %d\n", b.position.X, b.position.Y)
	for pieceRow := 0; pieceRow < len(b.shape); pieceRow++ {
		for pieceColumn := 0; pieceColumn < len(b.shape[pieceRow]); pieceColumn++ {
			if b.shape[pieceRow][pieceColumn] {
				b.data[b.position.Y+pieceRow][b.position.X+pieceColumn] = true
			}
		}
	}
}

func (b Board) Get() string {
	b.clear()
	b.paint()
	b.mux.Lock()
	defer b.mux.Unlock()
	output := ""
	for _, row := range b.data {
		for _, col := range row {
			if col {
				output += "🔳"
				continue
			}
			output += "🔲"
		}
		output += "\n"
	}

	return output
}
