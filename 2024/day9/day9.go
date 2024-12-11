package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadData() string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	var data string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = scanner.Text()
	}
	return data
}

func unCompress(data string) []int {
	id := 0
	unCompressed := []int{}

	for i := 0; i < len(data); i++ {
		ch := data[i] - '0'
		if i%2 == 0 {
			for j := 0; j < int(ch); j++ {
				unCompressed = append(unCompressed, id)
			}
			id++
		} else {
			for j := 0; j < int(ch); j++ {
				unCompressed = append(unCompressed, -1)
			}
		}
	}
	return unCompressed
}

func rearrange(d *[]int) {
	l := 0
	r := len(*d) - 1
	for l < r {
		for (*d)[l] != -1 {
			l++
		}
		for (*d)[r] == -1 {
			r--
		}
		if l < r {
			(*d)[l], (*d)[r] = (*d)[r], (*d)[l]
		}
	}
}
func rearrangeV2(d *[]int) {
	maxFileID := -1
	for _, v := range *d {
		if v > maxFileID {
			maxFileID = v
		}
	}

	// Iterate over files in decreasing order of file ID
	for fileID := maxFileID; fileID >= 0; fileID-- {
		// Find the current position and length of the file
		start, length := -1, 0
		for i, v := range *d {
			if v == fileID {
				if start == -1 {
					start = i
				}
				length++
			} else if start != -1 {
				break
			}
		}

		if start == -1 {
			continue // File not found
		}

		// Find the leftmost span of free space that can fit the file
		freeStart, freeLength := -1, 0
		for i, v := range *d {
			if v == -1 {
				if freeStart == -1 {
					freeStart = i
				}
				freeLength++
				if freeLength == length {
					break
				}
			} else {
				freeStart, freeLength = -1, 0
			}
		}

		// Move the file if a suitable span is found
		if freeLength == length && freeStart < start {
			for i := 0; i < length; i++ {
				(*d)[freeStart+i] = fileID
				(*d)[start+i] = -1
			}
		}
	}
}

func calcChecksum(data []int) int {
	checksum := 0
	for i := 0; i < len(data); i++ {
		if data[i] != -1 {
			checksum += i * data[i]
		}
	}
	return checksum
}

func main() {
	data := loadData()
	intData := unCompress(data)

	//rearrange(&intData)
	rearrangeV2(&intData)

	//Sfmt.Println(intData)
	fmt.Println(calcChecksum(intData))
}
