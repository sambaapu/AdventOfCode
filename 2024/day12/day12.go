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

func loadData() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	data := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for _, char := range line {
			row = append(row, string(char))
		}
		data = append(data, row)
	}
	return data
}

func dfs(r int, c int, data *[][]string, prev string, visited *map[Point]bool, perimeter *int) int {
	if r < 0 || r >= len(*data) || c < 0 || c >= len((*data)[0]) {
		return 0
	}
	if (*data)[r][c] != prev {
		return 0
	}
	if (*visited)[Point{r, c}] {
		return 0
	}
	(*visited)[Point{r, c}] = true
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	count := 0
	area := 1
	for _, d := range dir {
		nr, nc := r+d[0], c+d[1]
		if nr < 0 || nr >= len(*data) || nc < 0 || nc >= len((*data)[0]) || (*data)[nr][nc] != prev {
			count++
			continue
		}
		area += dfs(nr, nc, data, prev, visited, perimeter)
	}
	*perimeter += count
	return area
}

func main() {
	data := loadData()
	visited := make(map[Point]bool)
	res := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if !visited[Point{i, j}] {
				perimeter := 0
				area := dfs(i, j, &data, data[i][j], &visited, &perimeter)
				fmt.Println("letter:", data[i][j], "area: ", area, "perimeter: ", perimeter)
				res += area * perimeter
			}
		}
	}
	fmt.Println(res)

}
