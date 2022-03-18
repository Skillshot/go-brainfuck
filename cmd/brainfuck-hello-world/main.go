package main

import (
	"github.com/Skillshot/go-brainfuck/brainfuck"
)

func main() {
	//helloWorldProgram := "[->+<].>." // move value from first cell to the second
	helloWorldProgram := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

	brainfuck.Interpret(helloWorldProgram)
}
