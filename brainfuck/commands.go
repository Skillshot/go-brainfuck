package brainfuck

import "fmt"

type MoveForward struct{}

func (MoveForward) execute(memory *memory) {
	memory.pointer++
}

type MoveBackward struct{}

func (MoveBackward) execute(memory *memory) {
	memory.pointer--
}

// IncrementCellValue is a command that increments the current cell value.
type IncrementCellValue struct{}

func (IncrementCellValue) execute(memory *memory) {
	memory.cells[memory.pointer]++
}

type DecrementCellValue struct{}

func (DecrementCellValue) execute(memory *memory) {
	memory.cells[memory.pointer]--
}

type OutputCurrentCell struct{}

func (OutputCurrentCell) execute(memory *memory) {
	fmt.Print(string(memory.cells[memory.pointer]))
	//fmt.Print(memory.cells[memory.pointer])
}

type Loop struct {
	commands []command
}

func (l Loop) execute(m *memory) {
	for m.cells[m.pointer] != 0 {
		for _, c := range l.commands {
			c.execute(m)
		}
	}
}
