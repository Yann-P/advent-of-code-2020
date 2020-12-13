package main

func GetCharacterRight3Down1(lines []string, x int, y int) byte {
	newY := (y + 1) % len(lines)
	newX := (x + 3) % len(lines[newY])
	return lines[newY][newX]
}
