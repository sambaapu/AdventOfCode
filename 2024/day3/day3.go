package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	res := 0

	content := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text()
		//i := strings.Index(content, "do()")

	}
	j := 0
	for j != -1 {
		matches := []string{}
		j = strings.Index(content, "don't()")

		if j == -1 {
			matches = r.FindAllString(content, -1)
		} else {
			matches = r.FindAllString(content[:j], -1)
		}

		for _, match := range matches {
			var a, b int
			match = strings.TrimPrefix(match, "mul(")
			match = strings.TrimSuffix(match, ")")
			_, _ = fmt.Sscanf(match, "%d,%d", &a, &b)
			res += a * b
			//fmt.Println(a, b)
		}
		content = content[j+7:]
		j = strings.Index(content, "do()")
		if j == -1 {
			break
		}
		content = content[j+4:]
	}
	fmt.Println(res)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
