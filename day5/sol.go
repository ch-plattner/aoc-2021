package day1

import (
	"github.com/ch-plattner/aoc-2021/util"
	"strconv"
	"strings"
)

//0,9 -> 5,9
//8,0 -> 0,8
//9,4 -> 3,4
//2,2 -> 2,1
//7,0 -> 7,4
//6,4 -> 2,0
//0,9 -> 2,9
//3,4 -> 1,4
//0,0 -> 8,8
//5,5 -> 8,2

type Coord struct {
	X int
	Y int
}

type Line struct {
	Start Coord
	End Coord
}

// x -> (y -> val)
func one(inputFile string) (int, error) {
	lines := parseInputFile(inputFile, true)
	lineMap := make(map[int]map[int]int, 0)
	for _, line := range lines {
		// ok to move like this because we assume vertical/horizontal lines
		for x := line.Start.X; x <= line.End.X; x++ {
			for y := line.Start.Y; y <= line.End.Y; y++ {
				markSpot(x, y, lineMap)
			}
		}
	}
	return countLineMap(lineMap), nil
}

func two(inputFile string) (int, error) {
	lines := parseInputFile(inputFile, false)
	lineMap := make(map[int]map[int]int, 0)
	for _, line := range lines {
		xDiff := util.Abs(line.Start.X - line.End.X)
		yDiff := util.Abs(line.Start.Y - line.End.Y)
		if yDiff == 0 { // horizontal line
			for x := line.Start.X; x <= line.End.X; x++ {
				markSpot(x, line.Start.Y, lineMap)
			}
		} else if xDiff == 0 { // vertical line
			for y := line.Start.Y; y <= line.End.Y; y++ {
				markSpot(line.Start.X, y, lineMap)
			}
		} else { // diagonal line
			i := 0
			increment := line.Start.Y < line.End.Y
			// we will always move left->right in the X field
			// we may need to move up or down in the Y field
			for x := line.Start.X; x <= line.End.X; x++ {
				y := line.Start.Y + i
				markSpot(x, y, lineMap)
				if increment {
					i += 1
				} else {
					i -= 1
				}
			}
		}
	}
	return countLineMap(lineMap), nil
}

func markSpot(x, y int, lineMap map[int]map[int]int) {
	if lineMap[x] == nil {
		lineMap[x] = make(map[int]int, 0)
	}
	lineMap[x][y] = lineMap[x][y] + 1
}

func countLineMap(lineMap map[int]map[int]int) int {
	var res int
	for x := range lineMap {
		for y := range lineMap[x] {
			if lineMap[x][y] > 1 {
				res += 1
			}
		}
	}
	return res
}

func parseInputFile(file string, filter bool) []Line {
	rawLines, err := util.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var res []Line
	for _, rawLine := range rawLines {
		line := parseIntoLine(rawLine)
		if !filter || (line.Start.X == line.End.X || line.Start.Y == line.End.Y) {
			res = append(res, line)
		}
	}
	return res
}

func parseIntoLine(rawLine string) Line {
	rawCoords := strings.Split(rawLine, " -> ")
	if len(rawCoords) != 2 {
		panic(rawLine)
	}
	first := parseCoord(rawCoords[0])
	second := parseCoord(rawCoords[1])
	if first.X < second.X {
		return Line{
			Start: first,
			End: second,
		}
	}
	if first.X > second.X {
		return Line{
			Start: second,
			End: first,
		}
	}
	if first.Y < second.Y {
		return Line{
			Start: first,
			End: second,
		}
	}
	return Line{
		Start: second,
		End: first,
	}
}

func parseCoord(rawCoord string) Coord {
	splitRawCoord := strings.Split(rawCoord, ",")
	x, err := strconv.ParseInt(splitRawCoord[0], 10, 32)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseInt(splitRawCoord[1], 10, 32)
	if err != nil {
		panic(err)
	}
	return Coord{
		X: int(x),
		Y: int(y),
	}
}