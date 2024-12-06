package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	res1 := []int{}
	res2 := []int{}

	scanner := bufio.NewScanner(file)
	var i int = 0
	for scanner.Scan() {
		var curr_line string = scanner.Text()
		var r1, r2 int
		_, err := fmt.Sscanf(curr_line, "%d %d", &r1, &r2)
		res1 = append(res1, r1)
		res2 = append(res2, r2)
		i++
		if err != nil {
			fmt.Println("Error")
		}
	}
	sort.Ints(res1)
	sort.Ints(res2)
	var s float64 = 0
	for i = 0; i < 1000; i++ {
		s += math.Abs(float64(res1[i] - res2[i]))
	}
	fmt.Println(s)
}
func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var m1 map[int]int
	var m2 map[int]int
	m1 = make(map[int]int)
	m2 = make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var curr_line string = scanner.Text()

		var r1, r2 int
		_, err := fmt.Sscanf(curr_line, "%d %d", &r1, &r2)
		if err != nil {
			fmt.Println("Error")
		}

		m1[r1]++
		m2[r2]++
	}
	var s int = 0
	for key, val := range m1 {

		s += key * val * m2[key]
	}
	fmt.Println(s)
}
func main() {
	part1()
	part2()
}
