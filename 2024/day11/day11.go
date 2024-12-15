package main

import (
	"fmt"
	"math"
	"strconv"
)

type Pair struct {
	num  int
	time int
}

func numOfSplit(num int, time int, cache *map[Pair]int) int {
	if time == 0 {
		(*cache)[Pair{num, time}] = 1
		return 1
	}
	if (*cache)[Pair{num, time}] != 0 {
		return (*cache)[Pair{num, time}]
	}
	numDigits := len(strconv.Itoa(num))
	if num == 0 {
		(*cache)[Pair{num, time}] = numOfSplit(1, time-1, cache)
		return (*cache)[Pair{num, time}]
	} else if numDigits%2 == 0 {
		first := num / int(math.Pow10(numDigits/2))
		second := num % int(math.Pow10(numDigits/2))
		(*cache)[Pair{num, time}] = numOfSplit(first, time-1, cache) + numOfSplit(second, time-1, cache)
		return (*cache)[Pair{num, time}]
	} else {
		(*cache)[Pair{num, time}] = numOfSplit(num*2024, time-1, cache)
		return (*cache)[Pair{num, time}]
	}
}
func simulate(nums []int, time int) int {
	cache := make(map[Pair]int)
	res := 0
	for _, num := range nums {
		res += numOfSplit(num, time, &cache)
	}
	return res
}
func main() {
	res := simulate([]int{20, 82084, 1650, 3, 346355, 363, 7975858, 0}, 75)

	fmt.Println(res)
}
