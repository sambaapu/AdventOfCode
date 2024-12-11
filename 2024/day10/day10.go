package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	r int
	c int
}

func loadData() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	data := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, char := range line {
			if char == '.' {
				row = append(row, -1)
			} else {
				row = append(row, int(char-'0'))
			}
		}
		data = append(data, row)
	}
	return data
}

func dfs(r int, c int, prev int, data *[][]int, count *map[Point]bool,
	visited *map[Point]bool, path *[]Point, paths *[][]Point) {
	if r < 0 || r >= len(*data) || c < 0 || c >= len((*data)[0]) {
		return
	}
	if (*data)[r][c] == -1 {
		return
	}
	if (*data)[r][c]-prev != 1 {
		return
	}
	if (*visited)[Point{r, c}] {
		return
	}
	if (*data)[r][c] == 9 {
		(*count)[Point{r, c}] = true
		*path = append(*path, Point{r, c})
		newPath := make([]Point, len(*path))
		copy(newPath, *path)
		*paths = append(*paths, newPath)
		*path = (*path)[:len(*path)-1]
		return
	}
	(*visited)[Point{r, c}] = true
	*path = append(*path, Point{r, c})
	prev = (*data)[r][c]
	dfs(r+1, c, prev, data, count, visited, path, paths)
	dfs(r-1, c, prev, data, count, visited, path, paths)
	dfs(r, c+1, prev, data, count, visited, path, paths)
	dfs(r, c-1, prev, data, count, visited, path, paths)
	*path = (*path)[:len(*path)-1]
	(*visited)[Point{r, c}] = false
}

func uniqueCount(data [][]int) int {
	res := 0
	paths := [][]Point{}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if data[i][j] == 0 {
				count := make(map[Point]bool)
				visited := make(map[Point]bool)
				dfs(i, j, -1, &data, &count, &visited, &[]Point{}, &paths)
				//fmt.Println(count)
				res += len(count)
			}
		}
	}
	fmt.Println("unique paths: ", len(paths))
	return res
}

func main() {
	data := loadData()
	fmt.Println("reachable 9: ", uniqueCount(data))
}
