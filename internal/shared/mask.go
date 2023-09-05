package shared

type Tile bool
type MaskRow []Tile
type Mask []MaskRow

func (m Mask) Width() int {
	w := 0
	for _, row := range m {
		if len(row) > w {
			w = len(row)
		}
	}
	return w
}

func (m Mask) Height() int {
	return len(m)
}
