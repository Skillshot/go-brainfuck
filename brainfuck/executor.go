package brainfuck

func Interpret(program string) {
	// As now we have a loop command it makes sense to optimize the program
	// so we don't go through the string forward and backward.
	// In order to do so we can iterate through the string only once and store the commands in a list, then execute it.
	// So the complexity will be O(2*N).
	commands := compile(program)
	run(commands, &memory{})
}

func compile(program string) []command {
	var commandsStack = &[][]command{{}}
	for _, symbol := range program {
		instructionMap[symbol](commandsStack)
	}
	return (*commandsStack)[len(*commandsStack)-1]
}

func run(commands []command, memory *memory) {
	for _, command := range commands {
		command.execute(memory)
	}
}

type processInstruction func(commandsStack *[][]command)

var instructionMap = map[rune]processInstruction{
	'>': func(commandsStack *[][]command) {
		addCommand(commandsStack, MoveForward{})
	},
	'<': func(commandsStack *[][]command) {
		addCommand(commandsStack, MoveBackward{})
	},
	'+': func(commandsStack *[][]command) {
		addCommand(commandsStack, IncrementCellValue{})
	},
	'-': func(commandsStack *[][]command) {
		addCommand(commandsStack, DecrementCellValue{})
	},
	'.': func(commandsStack *[][]command) {
		addCommand(commandsStack, OutputCurrentCell{})
	},
	'[': func(commandsStack *[][]command) {
		*commandsStack = append(*commandsStack, []command{})
	},
	']': func(commandsStack *[][]command) {
		topElement := len(*commandsStack) - 1
		accumulatedLoopInnerCommands := (*commandsStack)[topElement]
		*commandsStack = (*commandsStack)[:topElement]
		addCommand(commandsStack, Loop{commands: accumulatedLoopInnerCommands})
	},
}

func addCommand(commandsStack *[][]command, command command) {
	cmdStack := *commandsStack
	topElement := len(cmdStack) - 1
	commands := cmdStack[topElement]
	cmdStack[topElement] = append(commands, command)
}

type command interface {
	execute(memory *memory)
}
