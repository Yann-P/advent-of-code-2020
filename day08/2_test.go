package main

import (
	"testing"
)

func TestProgramTerminates(t *testing.T) {
	ans, acc := ProgramTerminates(ParseCommands(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`))
	if ans != false {
		t.Error(1, ans)
	}

	ans, acc = ProgramTerminates(ParseCommands(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
nop -4
acc +6`))
	if ans != true {
		t.Error(1, ans)
	}

	if acc != 8 {
		t.Error(2, acc)
	}
}

func TestFindProgramThatTerminates(t *testing.T) {
	program := ParseCommands(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)
	idx := FindProgramThatTerminates(program, "jmp", "nop")

	if idx != 7 {
		t.Error(1, idx)
	}

	modifiedProgram := program
	modifiedProgram[idx].command = "nop"

	ans, acc := ProgramTerminates(modifiedProgram)

	if ans != true {
		t.Error(1, ans)
	}

	if acc != 8 {
		t.Error(2, acc)
	}
}
