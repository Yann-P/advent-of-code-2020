package main

import (
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	program := ParseCommands(string(content))

	println("Part 1")
	println(RunProgram(program))

	println("Part 2")
	cmdToFlipIndex := FindProgramThatTerminates(program, "jmp", "nop")
	modifiedProgram := program
	modifiedProgram[cmdToFlipIndex].command = "nop"
	terminates, acc := ProgramTerminates(modifiedProgram)
	println(terminates, acc)
}
