package brainfuck

import "testing"

func TestInterpret(t *testing.T) {
	type args struct {
		program string
	}
	tests := []struct {
		name string
		args args
	}{
		{"hello world", args{"++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Interpret(tt.args.program)
		})
	}
}
