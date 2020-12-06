// desc: https://adventofcode.com/2020/day/3

package main

import "fmt"

// Position asodasjd
type Position struct {
	X int
	Y int
}

func checkPosition(xStep int, yStep int, position Position, input []string, trees int) int {
	if position.Y == (len(input) - 1) {
		return trees
	}

	position.X = position.X + xStep
	if position.X > (len(input[0]) - 1) {
		position.X = position.X - (len(input[0]))
	}
	position.Y = position.Y + yStep
	if string(input[position.Y][position.X]) == "#" {
		trees++
	}
	return checkPosition(xStep, yStep, position, input, trees)
}

func partOne(input []string) int {
	trees := 0
	position := Position{0, 0}
	return checkPosition(3, 1, position, input, trees)
}

func partTwo(input []string) int {
	trees := 0
	position := Position{0, 0}
	return checkPosition(1, 1, position, input, trees) * checkPosition(3, 1, position, input, trees) * checkPosition(5, 1, position, input, trees) *
		checkPosition(7, 1, position, input, trees) * checkPosition(1, 2, position, input, trees)
}

func main() {
	givenInput := []string{
		"...............#...#..#...#....",
		"...#....##.....##...######..#..",
		"....#.....#.##..#...#..#.......",
		"...#..........#.#....#......#..",
		"....#.#...#.#.......#......##.#",
		"....#.#..........#.....#.##....",
		"##...#.#.##......#......#.#.#..",
		"#.#.#........#....#..#.#.......",
		"..#...##..#..#.......#....###..",
		".#.#..........#...#........#..#",
		".#..#......#....#.#...#...#.#.#",
		"..#.........#..##.....#.#.##.#.",
		".#......#...#....#.....#......#",
		"........#..##..##.........#..#.",
		".....#....###..#....##........#",
		".###...#..##..#.##......##...##",
		".#...#...#..#.#..#...##.....#..",
		".......#....#....#.#...#.......",
		".##.......#.##...#.#...#..#....",
		"#.#...#......#....#.......#....",
		"..###...............####...#.#.",
		".##.#....#......#..#...#.#..#.#",
		".............#.#.#......##.....",
		"#....#.#.#........#....##...#..",
		"...##....##...##..#...#........",
		"..##......#...#......#...###...",
		"...#...##......##.#.#..#.......",
		"#......#..#...#.#..#......#..##",
		".#..#..#........##....##.......",
		".#...........###.###.##...#....",
		"............#.#...........#.#..",
		"#...#........#...#...#..#.#.#.#",
		"...#.......#.#.#..#.#..........",
		"......#..#..#....##..##..##....",
		"........##...#......#..#.#.....",
		"..#.#.......#......#...........",
		"#.#.....#......##..........#.#.",
		"#.....###.#...#...#..#....#...#",
		".##.#...#............##.....#..",
		"###....#.#.....#.......##......",
		"##.......##.....#..#...#..##.##",
		"....#.##............###...#..##",
		".###..#...##.#..#..##..#.......",
		".##.##..####.....#.........#...",
		"....####..#...#....#.#...#.....",
		"..##....#..#......#...........#",
		"..........#......#..##.......#.",
		".................#.#.#........#",
		"#.......##.#...##.......##.##..",
		".#................#.#.....##..#",
		"......#..#............##.#.....",
		"...##............#.....#.....#.",
		"##.###.#.......#....#.........#",
		"......#.....#.#.#.#......#.....",
		"......#.##......#......##......",
		"..#....#...#..#.....#..#....#.#",
		".#.##.##.....#.......#.......#.",
		"...#..#.#......##....##..#.....",
		".#.....#......##...#..#.#....#.",
		"..#......#....#..#..###.#.#....",
		".....#........#.#...#..#.#.....",
		"....#.#.......#...##.....####..",
		"#..#..##...#...........#...#..#",
		".#..#...#.....#.....#.#.....#.#",
		"..##..###.....#...#.#.#.......#",
		"#..##.##......###..#......###..",
		"#..#...#.......#....#.#...#.#.#",
		"...........###..#...#..##....#.",
		".....#...........###.......#...",
		"##......#.......#......##.#..#.",
		"#.................#........#...",
		"#.#.............#......#...#..#",
		"......#.#....#....#....#....#.#",
		"..#...#....#..#....#....#..#...",
		"#....#......#..#...#..#....#.#.",
		"..#.....#..#...#...#.......#...",
		".#........###..##.#..#.........",
		".....##.#.....#..###..#........",
		"...#...#.###....######......#..",
		".###.#..#.....#.....#..#...#...",
		"##..#.#......#.........#...#..#",
		"###...##..###.......#..##.#.#.#",
		".#...................#..#...##.",
		".#...#...###...#.......#......#",
		"....#.....##....#...##.#...#.##",
		"..#....#......#...#..###.......",
		".........#..........##........#",
		"...#.........##..#....#..###...",
		"......#.##....#.#........#.#.##",
		"....#..#...#.............#....#",
		"#..#.....#.#.....#....#........",
		"....#..#...####.#.###..#.......",
		".....#....#....#..#..#.#.....#.",
		"#..##....#.....#.#.............",
		".##...#..#.......#...##.###....",
		"##......#...##....#......##....",
		"#......#.#......#.#..#......#..",
		".#...#......#............###..#",
		".#..#.#.....#.#.....#..#....#..",
		".#............#...#..#..#...##.",
		"...#...#....#.#.#....#....#....",
		"........#......###.#....##.##.#",
		"......#.#..................##..",
		"..#..#........#..##..........##",
		".#...#.#....#####.#.#..#.......",
		".....#..#.#..#.....#.#........#",
		"#.#..#..#.#..........#..#..#...",
		".##........#.#.......#........#",
		".....#....#..##...#......##....",
		"....#.##.##...##.#.........#.#.",
		"...#...#..#.#.......#.......#..",
		".................##..#.........",
		"..##.##....#...#.##.......#..#.",
		"....#...........#.....###....##",
		".#..................#..........",
		"....#.##....#......##.#..#.##.#",
		"#......#..#.#....##...####...#.",
		"#.....#.#.##...........#.##...#",
		".......#...##..#.........##....",
		".#...#..........##......#..#.#.",
		"#...#.#......#.....#........#..",
		"........#.....#.......##......#",
		".#..#...........#...#..#..#....",
		"......#.##......##.#......#..##",
		"......#....#..#................",
		"##.#.......#......#..#....##.##",
		"..#...###..#.......#...#.#....#",
		".....#...#.........#.....#..#.#",
		"..#.....#.........#..#...#.#...",
		".#.........##..#.#...#.....##..",
		"..........#.#.#...#....#....#..",
		"....#.###.#..#..#..#.#........#",
		"..#...##...##.#.#.....#.#..##..",
		".#..#......#.####..#.......#..#",
		"##.......#.....#.#.#..#...##..#",
		".##......##...##.........#..#..",
		"..#.##..#......#......#..#..#..",
		"..#..#.##..#........#....#.#...",
		"##.####...##.#..##.........#..#",
		".#.......#.#..#.....#.##.......",
		"........#.........#...........#",
		"..#...#.####.....##.##.#....#.#",
		".....#..##.#..###.##.#.#...#...",
		"#..##..##....#...#....#...#....",
		".###.#....#..#.......#......###",
		".#....##.....#.#........#...#.#",
		"#.#.#..#..#...#....#..#.....#..",
		"....##...#.##..#............#..",
		"##......###...#...#...#....#...",
		"#.#...#.#..#..##.##..##..#....#",
		".#...#.#....#..##.#..##..#.....",
		".............#..#..#.#.....#...",
		"#........#..........#....###...",
		"..#..#......#...#........#.....",
		".#..#.............#.........#..",
		"#.....#....##..#..#.#.#..#.....",
		"...#......#.........#.#....##.#",
		"..#.......##....###..#.........",
		".#.......#......#..............",
		".#...##.....##...###.#....#.#..",
		"......#....#.........#.......#.",
		"##.......#..##....###.#.....##.",
		".....#......####.....#......#..",
		"....#....#..###....#.....##...#",
		"...#...#.#........#.....#..#.##",
		"#..#....#.#...###...##.....##..",
		"#.....##...#.#............#....",
		"##....#.......#.#.#....#.#.###.",
		"#........#...#...####.#....#.#.",
		".##.#......#........#..#.....#.",
		"...##.#.......#...#.#.##..#...#",
		"......#.........##..#....#.....",
		".#......#...........#......#...",
		"......#.#.#.#..#.#....#....#..#",
		"##.................##...#.#....",
		"........#.........#..#..#...###",
		".#........#........#....#..#...",
		"....###.##.##......#.#...#....#",
		"#......#.#..............#......",
		".......#..#....#..##.........#.",
		"#............#....#............",
		"#...#.#.........##..###...##...",
		".#....#.#.#.....#..#..##.......",
		".............##.#.#.......#..#.",
		"#...#..##.#..###.....##..#.....",
		"...#.#.....#...#......##.#..#..",
		"#..........#.##.#...#...#.#...#",
		"...#.#...#.........#.#..#.#..#.",
		"#....#.................#.......",
		"..#..#.#.#..#.#..##...#.#......",
		"..#....#.#.#...#.....#...#.....",
		"#...#.###......#.......#..#..#.",
		"#.#.##.##..............#.#.#..#",
		"#..........#.#.........#.###...",
		"...........#.......#.#..#......",
		"....#.#..#....#....#....##.....",
		"#..#...##########.........#.#..",
		"..#.............##........#.#..",
		"#.#.##......#...#.#....##..##..",
		"..##..#.#.#....#......#.#..#.#.",
		".#.#...#................#......",
		"#...#...#.....##.#...#.......#.",
		".....##.......#...#......#.##..",
		"....#.##..........#.#.##....##.",
		"...........##........#.#.#.#...",
		"..#...........###.#....#..#....",
		"..#.#...#...#.#........#.....#.",
		".##......##.....#........#.....",
		"....#.......#....#...#.##...#..",
		".#.....##.....#...##...........",
		"#....#.##.##...#...###.#####...",
		"..#.#......#.#.......#....#..#.",
		"..........#...#...........#....",
		".........#..#...#...#......#.##",
		".#...#.#...#.###.##..#........#",
		"#....#.....#.......##....#.....",
		"#.....#..#.....#...##.......#..",
		"#.#.#.#.............#.....#...#",
		"...#.....#.....#...............",
		"..##.###.#.#........#.........#",
		"...##..#.##..#....#.#...#.#....",
		"...##...#...##.#.........#.#..#",
		".###......#....#.........#.#...",
		"###.#.#...#.......#...#.....#.#",
		"...#.#.......#.....####........",
		"......#..#.....###.#....#..#...",
		"..####...#...#..#......#...#...",
		"#..............##....#......##.",
		"..##..###...##..#.#.........#..",
		"#..#.#...#.......#.........##..",
		"..##....#......#...#..##.......",
		"..#.#..#..###.....#.....#...###",
		".#..#.##.....##.#.#.#........#.",
		"..#.#.........#................",
		"..#...........#.#...##.#...#..#",
		".....#........#..#.....#....#..",
		"#.#....#...........##.....#..##",
		"##.......#.....#.....#.#......#",
		".##............####.#........##",
		"....##.#.....#......#....#...#.",
		".#.#...##.#..##..#..........#..",
		"..........................#....",
		"##..##..#..#.#....#.....#......",
		"...#.#........#.#.##.#.....#..#",
		"#..#....#...#..#...#........#.#",
		"#.......#####...#..#..#.#......",
		"#.##....#......#......#..###...",
		"..#.......#...........#.....##.",
		"#...#....#..#......##...#......",
		"...##.#..##........#..###......",
		"##.#...........#.............##",
		"#.....#..#.....#.....#.........",
		".#..........#..#......###......",
		"...#...#..##.......#...#...#.#.",
		"..####......#.....#..#.#......#",
		"....#..#.....#.#...............",
		".#.......#....#.....#..##....#.",
		".....#.........#..........##...",
		"#...........#..#.....#........#",
		"............#..##...#...#.#....",
		"##.............####...#.....#..",
		".....#......#....##.#.....##...",
		"...........#.....#.#..#.#......",
		".#.#......#....#.............##",
		"##...#.#.......##........#.....",
		"#..#.....#.#.......####...#..#.",
		"....#.#...#....#.###..#..#.#...",
		".....#.#..#......#.##.#####.#..",
		".....#....##..........#......#.",
		"#.......#........##.......##...",
		"#...#..#..##.#....#...#...#....",
		"...#..##..#...#..........#.....",
		"#..#....###.....#......##......",
		"...###......#.....#....#.......",
		"#.#...#.#..###.####.......#.#.#",
		"...#......#.#..##..#.....#.....",
		"#.#............#.....##.#......",
		"#..#......##...##..#...#.#..###",
		"#.....#...#..#..#....#..###....",
		"####..###....#..##....#..#.....",
		"..##.#.....#.......##....#.#.#.",
		"##..#.#.......#.#...##.#..#.#..",
		"..#.#.#.##.#.#.#...........#...",
		"...#.##.....#....#..#.#..#.....",
		"...#..#.........#..........#..#",
		"#...#..#.....#.#.#.............",
		"##.#....##.##...#...#..#..#..#.",
		"....####..##..##.#..#...##.....",
		"##.....##.#.#...#.#.......###..",
		"#..#.#....#......#.......##...#",
		"#.#...............#..#..#......",
		".....##...##..#........#......#",
		".#..#............##......#....#",
		"......#.#..#..##.#......#.....#",
		"..#.#.............#...#......##",
		"....#.#..#..#...##...#..##.....",
		"#.#.............#...#.#..#....#",
		"#..#..#.##....#....#...#.......",
		"....#..#..........#.....##...#.",
		"..#####.......#...#..........#.",
		"....#......##.......#..##.#.#.#",
		"#...#.#.....#....#....##.......",
		"..##.#.#..#.#...#.....##.....#.",
		"#.#..#....#............#..#.#..",
		"...#.##..##..#.#...###......#.#",
		"......##.......#....#......###.",
		"....#..##.......#..#.#....#..#.",
		"#............#..........##..#.#",
		"..#.....#..#.#..##..#....##.#..",
		".....#.....#....##.#.#......#.#",
		"...####.......###..#...#.#....#",
		".#.##.....##.....##..#..#......",
		"...#..###..##..................",
		"##..##.....#.#..#..#.#........#",
		".#.#........##.........#....#..",
		"........#......###....#...#....",
		"......#...........#...........#",
		".#...###...#.........#.#...#..#",
		".....#..........#......##....#.",
		".#.#...#........##.......#.#...",
		".....##.....##....#...#...#..#.",
		"..#.....................##...##",
		"#..#.#......##...##..#......#.."}

	fmt.Println(partOne(givenInput))
	fmt.Println(partTwo(givenInput))
}
