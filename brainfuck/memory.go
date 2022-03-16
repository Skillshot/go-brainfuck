package brainfuck

type memory struct {
	cells   [200]byte
	pointer int
}

func newMemory(pointer int, initCells map[int]byte) *memory {
	cells := [200]byte{}
	for i, val := range initCells {
		cells[i] = val
	}
	return &memory{pointer: pointer, cells: cells}
}
