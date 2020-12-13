package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	takenSeats := [0b1111111111]bool{false}

	highest := 0

	for _, code := range lines {
		if len(code) != 10 {
			continue
		}
		seatId := ExtractSeatId(code)

		takenSeats[seatId] = true
		if seatId > highest {
			highest = seatId
		}
	}

	println("Part 1")
	println(highest)

	freeSeatIds := []int{}

	for seatId := 0; seatId < 0b1111111111; /* 7 bits row, 3 bits col */ seatId++ {
		rowNumber := ExtractRowNumberFromSeatId(seatId)
		if takenSeats[seatId] != true && rowNumber != 0b1111111 && rowNumber != 0 {
			freeSeatIds = append(freeSeatIds, seatId)
		}
	}

	for i, freeSeatId := range freeSeatIds {
		if i == 0 || i == len(freeSeatIds)-1 {
			continue
		}
		if freeSeatIds[i-1] == freeSeatId-1 || freeSeatIds[i+1] == freeSeatId+1 {
			continue
		}
		println("Part 2")
		println(freeSeatId)
	}

}
