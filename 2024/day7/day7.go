package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func loadData() ([]int, [][]int) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	targets := []int{}
	data := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		inputs := strings.Split(line, ":")

		target, _ := strconv.Atoi(inputs[0])
		targets = append(targets, target)

		inputs[1] = strings.TrimSpace(inputs[1])
		dataLine := strings.Split(inputs[1], " ")
		currData := []int{}
		for _, v := range dataLine {
			num, _ := strconv.Atoi(v)
			currData = append(currData, num)
		}
		data = append(data, currData)
	}
	return targets, data
}
func backtrack(arr []int, i int, curr int, target int) bool {
	//fmt.Println("i", i, "curr", curr)
	if i == len(arr) {
		return curr == target
	}
	return backtrack(arr, i+1, curr+arr[i], target) || backtrack(arr, i+1, curr*arr[i], target)
}

func backtrackV2(arr []int, i int, curr int, target int) bool {
	//fmt.Println("i", i, "curr", curr)
	if i == len(arr) {
		return curr == target
	}
	n := len(strconv.Itoa(arr[i]))
	concatVal := curr*int(math.Pow(10, float64(n))) + arr[i]
	return backtrackV2(arr, i+1, curr+arr[i], target) ||
		backtrackV2(arr, i+1, curr*arr[i], target) ||
		backtrackV2(arr, i+1, concatVal, target)
}

func calculation(targets []int, data [][]int) (int, int) {
	res1 := 0
	res2 := 0
	for i, v := range targets {
		if backtrack(data[i], 1, data[i][0], v) {
			res1 += v
		}
		if backtrackV2(data[i], 1, data[i][0], v) {
			res2 += v
		}
	}
	return res1, res2
}
func main() {
	//flag := backtrack([]int{1, 2, 3}, 1, 1, 6)
	targets, data := loadData()
	res1, res2 := calculation(targets, data)
	fmt.Println("Part1: ", res1)
	fmt.Println("Part2: ", res2)
}
