package main

func GetCharacterRightDown(lines []string, x int, y int, dx int, dy int) byte {
	newY := (y + dy) % len(lines)
	newX := (x + dx) % len(lines[newY])
	return lines[newY][newX]
}

func ComputeTreesEncountered(lines []string, dx int, dy int) int {
	x := 0
	y := 0
	treesHit := 0

	for y < len(lines)-dy-1 {
		if GetCharacterRightDown(lines, x, y, dx, dy) == '#' {
			treesHit++
		}
		x += dx
		y += dy
	}

	return treesHit
}
