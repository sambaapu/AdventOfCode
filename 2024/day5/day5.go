package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func loadData() ([][]int, map[int][]int) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	order := [][]int{}
	graph := make(map[int][]int)
	scanner := bufio.NewScanner(file)
	flag := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			flag = true
			continue
		}

		var a, b int
		if flag {
			s := strings.Split(line, ",")
			intVec := []int{}
			for i := 0; i < len(s); i++ {
				num, _ := strconv.Atoi(s[i])
				intVec = append(intVec, num)
			}
			order = append(order, intVec)
		} else {
			_, _ = fmt.Sscanf(line, "%d|%d", &a, &b)
			if graph[a] == nil {
				graph[a] = []int{}
			}
			graph[a] = append(graph[a], b)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return order, graph
}

func isOrdered(orderVec []int, graph map[int][]int) bool {
	for i := 0; i < len(orderVec)-1; i++ {
		prev := orderVec[i]
		for j := i + 1; j < len(orderVec); j++ {
			next := orderVec[j]
			if graph[next] != nil {
				for _, v := range graph[next] {
					if v == prev {
						//fmt.Println(prev, next)
						return false
					}
				}
			}
		}
	}
	return true
}

func sortVec(orderVec []int, graph map[int][]int) []int {
	sort.Slice(orderVec, func(i, j int) bool {
		prev, next := orderVec[i], orderVec[j]
		if graph[next] != nil {
			for _, v := range graph[next] {
				if v == prev {
					return false
				}
			}
		}
		return true
	})
	return orderVec
}
func main() {
	order, graph := loadData()
	res1 := 0
	res2 := 0
	for _, orderVec := range order {
		if isOrdered(orderVec, graph) {
			res1 += orderVec[len(orderVec)/2]
		} else {
			orderVec = sortVec(orderVec, graph)
			res2 += orderVec[len(orderVec)/2]
		}
	}
	fmt.Println(res1)
	fmt.Println(res2)
}
