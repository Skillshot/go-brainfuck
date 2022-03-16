package brainfuck

import (
	"reflect"
	"testing"
)

func TestLoop_execute(t *testing.T) {
	type fields struct {
		commands []command
	}
	type args struct {
		memory *memory
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected memory
	}{
		{
			"skips executing commands if memory current cell value is zero",
			fields{[]command{MoveForward{}, IncrementCellValue{}}},
			args{memory: &memory{}},
			memory{},
		},
		{
			"executes all commands if memory current cell value is non-zero",
			fields{[]command{
				MoveForward{}, IncrementCellValue{}, IncrementCellValue{}, MoveBackward{}, DecrementCellValue{},
			}},
			args{memory: newMemory(0, map[int]byte{0: 1})},
			*newMemory(0, map[int]byte{0: 0, 1: 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Loop{
				commands: tt.fields.commands,
			}
			l.execute(tt.args.memory)
			if !reflect.DeepEqual(*tt.args.memory, tt.expected) {
				t.Errorf("Expected %v \n but have %v", tt.expected, *tt.args.memory)
			}
		})
	}
}
