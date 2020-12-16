package main

import (
	"strconv"
	"strings"
)

type Command struct {
	command  string
	argument int
}

func ParseCommand(command string) Command {
	splitted := strings.Split(command, " ")
	argument, err := strconv.Atoi(splitted[1])

	if err != nil {
		panic(err)
	}

	return Command{
		command:  splitted[0],
		argument: argument,
	}
}

func RunProgram(program []Command) int {
	pointer := 0
	accumulator := 0
	visitedCommands := make(map[int]bool)

	for pointer <= len(program) && !visitedCommands[pointer] {
		visitedCommands[pointer] = true

		command := program[pointer]
		switch command.command {
		case "nop":
			pointer++
			break
		case "jmp":
			pointer += command.argument
			break
		case "acc":
			accumulator += command.argument
			pointer++
			break
		default:
			panic("Unknown command")
		}

	}

	return accumulator
}
