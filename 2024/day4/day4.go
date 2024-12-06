package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func loadData() [][]string {
	file, err := os.Open("./day4/input.txt")
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
		//fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return charMatrix
}

func dfs(charMatrix [][]string, r int, c int, dir string, curr string) int {
	ROW := len(charMatrix)
	COL := len(charMatrix[0])
	if r < 0 || c < 0 || r >= ROW || c >= COL {
		return 0
	}
	curr += charMatrix[r][c]
	if curr == "XMAS" {
		return 1
	}
	if curr != "XMAS"[:len(curr)] {
		return 0
	}
	switch dir {
	case "top":
		return dfs(charMatrix, r-1, c, "top", curr)
	case "down":
		return dfs(charMatrix, r+1, c, "down", curr)
	case "left":
		return dfs(charMatrix, r, c-1, "left", curr)
	case "right":
		return dfs(charMatrix, r, c+1, "right", curr)
	case "downright":
		return dfs(charMatrix, r+1, c+1, "downright", curr)
	case "downleft":
		return dfs(charMatrix, r+1, c-1, "downleft", curr)
	case "upright":
		return dfs(charMatrix, r-1, c+1, "upright", curr)
	case "upleft":
		return dfs(charMatrix, r-1, c-1, "upleft", curr)
	}
	return 0
}
func part1(charMatrix [][]string) {
	ROW := len(charMatrix)
	COL := len(charMatrix[0])
	res := 0
	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			if charMatrix[i][j] == "X" {
				res += dfs(charMatrix, i, j, "top", "")
				res += dfs(charMatrix, i, j, "down", "")
				res += dfs(charMatrix, i, j, "left", "")
				res += dfs(charMatrix, i, j, "right", "")
				res += dfs(charMatrix, i, j, "downright", "")
				res += dfs(charMatrix, i, j, "downleft", "")
				res += dfs(charMatrix, i, j, "upright", "")
				res += dfs(charMatrix, i, j, "upleft", "")
			}
		}
	}
	fmt.Println(res)

}

func main() {
	charMatrix := loadData()
	part1(charMatrix)
}
