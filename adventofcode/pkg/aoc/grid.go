package aoc

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

func MakeRuneGrid(width, height int, empty rune) [][]rune {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = empty
		}
	}

	return grid
}

func RenderRuneGrid(grid [][]rune) string {
	s := ""
	for _, l := range grid {
		s += string(l) + "\n"
	}

	return s
}

func HashRuneGrid(grid [][]rune) string {
	s := ""
	for _, l := range grid {
		s += string(l)
	}

	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

func CountRuneGridOccurrences(grid [][]rune, r rune) int {
	count := 0
	for _, l := range grid {
		for _, c := range l {
			if c == r {
				count += 1
			}
		}
	}

	return count
}

func RenderStringGrid(grid [][]string) string {
	s := ""
	for _, l := range grid {
		s += strings.Join(l, " ") + "\n"
	}

	return s
}
func RenderIntGridS(grid [][]int, sep string) string {
	s := ""
	for _, l := range grid {
		s += IntArrayAsString(l, sep) + "\n"
	}

	return s
}
