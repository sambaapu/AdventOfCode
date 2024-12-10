package main

import (
	"bufio"
	"fmt"
	"os"
)

var X int = 0
var Y int = 0

type coordinate struct {
	x int
	y int
}

func loadData() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
func getPosMap(data []string) map[int][]coordinate {
	posMap := make(map[int][]coordinate)
	for i, v := range data {
		for k := range v {
			if v[k] == '.' {
				continue
			}
			pos := coordinate{i, k}
			antenna := v[k]
			posMap[int(antenna)] = append(posMap[int(antenna)], pos)
		}
	}
	return posMap
}
func myMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func myMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func myAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func inBound(c coordinate) bool {
	if c.x < 0 || c.y < 0 || c.x >= X || c.y >= Y {
		return false
	}
	return true
}
func getAntinodes(posMap map[int][]coordinate) map[coordinate]bool {
	antinodes := make(map[coordinate]bool)
	for _, v := range posMap {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				xdiff := myAbs(v[i].x - v[j].x)
				ydiff := myAbs(v[i].y - v[j].y)
				a1 := coordinate{-1, -1}
				a2 := coordinate{-1, -1}
				if (myMin(v[i].x, v[j].x) == v[i].x && myMin(v[i].y, v[j].y) == v[i].y) ||
					(myMin(v[i].x, v[j].x) == v[j].x && myMin(v[i].y, v[j].y) == v[j].y) {
					a1 = coordinate{myMin(v[i].x, v[j].x) - xdiff, myMin(v[i].y, v[j].y) - ydiff}
					a2 = coordinate{myMax(v[i].x, v[j].x) + xdiff, myMax(v[i].y, v[j].y) + ydiff}
				} else {
					a1 = coordinate{myMin(v[i].x, v[j].x) - xdiff, myMax(v[i].y, v[j].y) + ydiff}
					a2 = coordinate{myMax(v[i].x, v[j].x) + xdiff, myMin(v[i].y, v[j].y) - ydiff}
				}
				if inBound(a1) {
					antinodes[a1] = true
				}
				if inBound(a2) {
					antinodes[a2] = true
				}
			}

		}
	}
	return antinodes
}

func getHarmonicAntinodes(posMap map[int][]coordinate) map[coordinate]bool {
	antinodes := make(map[coordinate]bool)
	for _, v := range posMap {
		for i := 0; i < len(v); i++ {
			antinodes[v[i]] = true // add the node itself
			for j := i + 1; j < len(v); j++ {
				xdiff := myAbs(v[i].x - v[j].x)
				ydiff := myAbs(v[i].y - v[j].y)
				a1 := coordinate{-1, -1}
				a2 := coordinate{-1, -1}
				if (myMin(v[i].x, v[j].x) == v[i].x && myMin(v[i].y, v[j].y) == v[i].y) ||
					(myMin(v[i].x, v[j].x) == v[j].x && myMin(v[i].y, v[j].y) == v[j].y) {
					a1 = coordinate{myMin(v[i].x, v[j].x) - xdiff, myMin(v[i].y, v[j].y) - ydiff}
					for inBound(a1) {
						antinodes[a1] = true
						a1.x -= xdiff
						a1.y -= ydiff
					}
					a2 = coordinate{myMax(v[i].x, v[j].x) + xdiff, myMax(v[i].y, v[j].y) + ydiff}
					for inBound(a2) {
						antinodes[a2] = true
						a2.x += xdiff
						a2.y += ydiff
					}
				} else {
					a1 = coordinate{myMin(v[i].x, v[j].x) - xdiff, myMax(v[i].y, v[j].y) + ydiff}
					for inBound(a1) {
						antinodes[a1] = true
						a1.x -= xdiff
						a1.y += ydiff
					}
					a2 = coordinate{myMax(v[i].x, v[j].x) + xdiff, myMin(v[i].y, v[j].y) - ydiff}
					for inBound(a2) {
						antinodes[a2] = true
						a2.x += xdiff
						a2.y -= ydiff
					}
				}
			}

		}
	}
	return antinodes
}

func main() {
	data := loadData()
	X = len(data)
	Y = len(data[0])
	posMap := getPosMap(data)
	antinodes := getAntinodes(posMap)
	harmonicAntinodes := getHarmonicAntinodes(posMap)
	fmt.Println(len(antinodes))
	fmt.Println(len(harmonicAntinodes))
}
