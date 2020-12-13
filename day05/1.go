package main

func ParseRowNumber(code string) int {
	res := 0

	for i, letter := range code { // 7 bits
		if letter == 'B' {
			res |= (1 << (6 - i))
		}
	}

	return res
}

func ParseColNumber(code string) int {
	res := 0
	for i, letter := range code { // 3 bits
		if letter == 'R' {
			res |= (1 << (2 - i))
		}
	}
	return res
}

func ExtractSeatId(code string) int {
	row := ParseRowNumber(code[:7])
	col := ParseColNumber(code[7:10])
	return row*8 + col
}

func ExtractRowNumberFromSeatId(seatId int) int {
	return seatId >> 3
}
