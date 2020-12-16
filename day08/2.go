package main

import (
	"strings"
)

type State = struct {
	pointer     int
	accumulator int
	terminated  bool
}

func RunCommand(state State, command Command) State {
	switch command.command {
	case "nop":
		state.pointer++
		break
	case "jmp":
		state.pointer += command.argument
		break
	case "acc":
		state.accumulator += command.argument
		state.pointer++
		break
	default:
		panic("Unknown command")
	}
	return state
}

func ParseCommands(programAsString string) []Command {
	commandsAsString := strings.Split(programAsString, "\n")
	program := []Command{}

	for _, commandAsString := range commandsAsString {
		if commandAsString != "" {
			program = append(program, ParseCommand(commandAsString))
		}
	}

	return program
}

func ProgramTerminates(program []Command) (bool, int) {

	state := State{
		pointer:     0,
		accumulator: 0,
	}

	visitedCommands := make(map[int]bool)

	for state.pointer < len(program) {
		visitedCommands[state.pointer] = true

		state = RunCommand(state, program[state.pointer])

		if visitedCommands[state.pointer] == true {
			return false, -1
		}

	}

	return state.pointer == len(program), state.accumulator
}

func FindProgramThatTerminates(program []Command, flipWhat string, flipToWhat string) int {
	flippableCommandIndices := []int{}

	for i, command := range program {
		if command.command == flipWhat {
			flippableCommandIndices = append(flippableCommandIndices, i)
		}
	}

	for _, index := range flippableCommandIndices {
		modifiedProgram := []Command{}

		// Copy
		for _, command := range program {
			modifiedProgram = append(modifiedProgram, command)
		}

		modifiedProgram[index].command = flipToWhat
		teminates, _ := ProgramTerminates(modifiedProgram)
		if teminates {
			return index
		}
	}

	return -1
}
