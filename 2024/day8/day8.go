package main

import (
	"bufio"
	"fmt"
	"os"
)

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
func main() {
	data := loadData()
	fmt.Println(data)
	//fmt.Println(data[0][0])
}
