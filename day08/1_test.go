package main

import (
	"testing"
)

func TestParseCommand(t *testing.T) {
	command := ParseCommand("jmp +50")
	if command.command != "jmp" {
		t.Error("1")
	}

	if command.argument != 50 {
		t.Error("2")
	}

	command = ParseCommand("jmp -50")

	if command.argument != -50 {
		t.Error("3")
	}
}

func TestRunProgram(t *testing.T) {
	program := []Command{
		ParseCommand("nop +0"),
		ParseCommand("acc +1"),
		ParseCommand("jmp +4"),
		ParseCommand("acc +3"),
		ParseCommand("jmp -3"),
		ParseCommand("acc -99"),
		ParseCommand("acc +1"),
		ParseCommand("jmp -4"),
		ParseCommand("acc +6"),
	}

	ans := RunProgram(program)

	if ans != 5 {
		t.Error(ans)
	}
}
