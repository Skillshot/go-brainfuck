package main

import (
	"fmt"
)

type Memory struct {
	cells   [200]byte
	pointer int
}

func NewMemory(pointer int, initCells map[int]byte) *Memory {
	cells := [200]byte{}
	for i, val := range initCells {
		cells[i] = val
	}
	return &Memory{pointer: pointer, cells: cells}
}

func main() {
	//moveFirstNumberToNextCell := "[->+<].>."
	moveFirstNumberToNextCell := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

	Interpret(moveFirstNumberToNextCell)
}

type ProcessInstruction func(commandsStack *[][]Command)

type MoveForward struct{}

func (MoveForward) execute(memory *Memory) {
	memory.pointer++
}

type MoveBackward struct{}

func (MoveBackward) execute(memory *Memory) {
	memory.pointer--
}

type IncrementCellValue struct{}

func (IncrementCellValue) execute(memory *Memory) {
	memory.cells[memory.pointer]++
}

type DecrementCellValue struct{}

func (DecrementCellValue) execute(memory *Memory) {
	memory.cells[memory.pointer]--
}

type OutputCurrentCell struct{}

func (OutputCurrentCell) execute(memory *Memory) {
	fmt.Print(string(memory.cells[memory.pointer]))
	//fmt.Print(memory.cells[memory.pointer])
}

type Loop struct {
	commands []Command
}

func (l Loop) execute(memory *Memory) {
	for memory.cells[memory.pointer] != 0 {
		for _, c := range l.commands {
			c.execute(memory)
		}
	}
}

var instructionMap = map[rune]ProcessInstruction{
	'>': func(commandsStack *[][]Command) {
		addCommand(commandsStack, MoveForward{})
	},
	'<': func(commandsStack *[][]Command) {
		addCommand(commandsStack, MoveBackward{})
	},
	'+': func(commandsStack *[][]Command) {
		addCommand(commandsStack, IncrementCellValue{})
	},
	'-': func(commandsStack *[][]Command) {
		addCommand(commandsStack, DecrementCellValue{})
	},
	'.': func(commandsStack *[][]Command) {
		addCommand(commandsStack, OutputCurrentCell{})
	},
	'[': func(commandsStack *[][]Command) {
		*commandsStack = append(*commandsStack, []Command{})
	},
	']': func(commandsStack *[][]Command) {
		topElement := len(*commandsStack) - 1
		accumulatedLoopInnerCommands := (*commandsStack)[topElement]
		*commandsStack = (*commandsStack)[:topElement]
		addCommand(commandsStack, Loop{commands: accumulatedLoopInnerCommands})
	},
}

func addCommand(commandsStack *[][]Command, command Command) {
	cmdStack := *commandsStack
	topElement := len(cmdStack) - 1
	commands := cmdStack[topElement]
	cmdStack[topElement] = append(commands, command)
}

func Interpret(program string) {
	// As now we have a loop command it makes sense to optimize the program
	// so we don't go through the string forward and backward.
	// In order to do so we can iterate through the string only once and store the commands in a list, then execute it.
	// So the complexity will be O(2*N).
	commands := compile(program)
	run(commands, &Memory{})
}

func run(commands []Command, memory *Memory) {
	for _, command := range commands {
		command.execute(memory)
	}
}

type Command interface {
	execute(memory *Memory)
}

func compile(program string) []Command {
	var commandsStack = &[][]Command{{}}
	for _, symbol := range program {
		instructionMap[symbol](commandsStack)
	}
	return (*commandsStack)[len(*commandsStack)-1]
}
