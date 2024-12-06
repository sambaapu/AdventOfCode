package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafe(curr_vec []int) bool {
	prev := curr_vec[0]
	dir := true
	if curr_vec[1]-prev < 0 {
		dir = false
	}
	for i := 1; i < len(curr_vec); i++ {
		c := curr_vec[i]
		dif := c - prev
		//fmt.Println(dif)
		if dir && (dif < 0) {
			return false
		}
		if (!dir) && (dif >= 0) {
			return false
		}
		if !dir {
			dif = -dif
		}
		if (dif > 3) || (dif < 1) {
			return false
		}
		prev = c
	}
	return true
}
func modifyVec(vec []int, i int) []int {
	curr_vec := make([]int, 0)
	curr_vec = append(curr_vec, vec[:i]...)
	curr_vec = append(curr_vec, vec[i+1:]...)
	return curr_vec
}
func modifiedVecSafe(vec []int) bool {
	for i := 0; i < len(vec); i++ {
		curr_vec := modifyVec(vec, i)
		//fmt.Println(curr_vec)
		if isSafe(curr_vec) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		var curr_line string = scanner.Text()
		strSlices := strings.Split(curr_line, " ")
		curr_vec := []int{}
		for i := 0; i < len(strSlices); i++ {
			curr_int, _ := strconv.Atoi(strSlices[i])
			curr_vec = append(curr_vec, curr_int)
		}
		if isSafe(curr_vec) {
			count += 1
		} else if modifiedVecSafe(curr_vec) {
			count += 1
		}
	}
	//fmt.Println(isSafe([]int{7, 6, 4, 2, 1}))
	fmt.Println(modifiedVecSafe([]int{1, 3, 2, 4, 5}))
	fmt.Println(count)
}
