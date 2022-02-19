package main

import (
	"fmt"
)

func main() {
	outputAB := "+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++." +
		">++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++."

	Interpret(outputAB)
}

type Memory struct {
	cells   [200]byte
	pointer int
}

// first try is to just define a function type which accepts the memory and executes the command

// first error that student most probably makes is passing Memory as a value not a pointer

type ExecuteCommand func(memory *Memory)

var commandMap = map[rune]ExecuteCommand{
	'>': func(memory *Memory) { memory.pointer++ },
	'<': func(memory *Memory) { memory.pointer-- },
	'+': func(memory *Memory) { memory.cells[memory.pointer]++ },
	'-': func(memory *Memory) { memory.cells[memory.pointer]-- },
	'.': func(memory *Memory) { fmt.Print(string(memory.cells[memory.pointer])) },
}

func Interpret(program string) {
	m := Memory{}
	for _, instructionSymbol := range program {
		commandMap[instructionSymbol](&m)
	}
}
