package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func loadData() [][]string {
	file, err := os.Open("./input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	charMatrix := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for _, ch := range line {
			row = append(row, string(ch))
		}
		charMatrix = append(charMatrix, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return charMatrix
}
func guardWalk(charMatrix [][]string, r int, c int, dir string) {
	ROW := len(charMatrix)
	COL := len(charMatrix[0])
	if r < 0 || c < 0 || r >= ROW || c >= COL {
		return
	}
	// if charMatrix[r][c] == "X" {
	// 	return
	// }
	charMatrix[r][c] = "X"
	switch dir {
	case "top":
		if charMatrix[r-1][c] == "#" {
			dir = "right"
			guardWalk(charMatrix, r, c+1, dir)
		} else {
			guardWalk(charMatrix, r-1, c, dir)
		}
	case "down":
		if charMatrix[r+1][c] == "#" {
			dir = "left"
			guardWalk(charMatrix, r, c-1, dir)
		} else {
			guardWalk(charMatrix, r+1, c, dir)
		}
	case "left":
		if charMatrix[r][c-1] == "#" {
			dir = "top"
			guardWalk(charMatrix, r-1, c, dir)
		} else {
			guardWalk(charMatrix, r, c-1, dir)
		}
	case "right":
		if charMatrix[r][c+1] == "#" {
			dir = "down"
			guardWalk(charMatrix, r+1, c, dir)
		} else {
			guardWalk(charMatrix, r, c+1, dir)
		}
	}
}

func main() {
	charMatrix := loadData()
	ROW := len(charMatrix)
	COL := len(charMatrix[0])
	res := 0
	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			if charMatrix[i][j] == "^" {
				guardWalk(charMatrix, i, j, "top")
			}
		}
	}
	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			if charMatrix[i][j] == "X" {
				res++
			}
		}
	}
	for i := 0; i < ROW; i++ {
		fmt.Println(charMatrix[i])
	}
	fmt.Println("Part1 result: ", res)
}
