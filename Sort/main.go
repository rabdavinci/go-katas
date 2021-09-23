package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ai := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ai), func(i, j int) { ai[i], ai[j] = ai[j], ai[i] })
	fmt.Println(ai)
	ai = sort(ai)

	fmt.Println(ai)
}

func sort(ai []int) []int {
	if len(ai) == 1 {
		return ai
	}
	pointer := ai[0]
	fmt.Println(pointer)
	var leftside, rightside, middle []int
	middle = append(middle, ai[0])
	for i := 1; i < len(ai); i++ {
		if ai[i] < pointer {
			leftside = append(leftside, ai[i])
		} else if ai[i] == pointer {
			middle = append(middle, ai[i])
		} else {
			rightside = append(rightside, ai[i])
		}
	}
	if len(leftside) != 1 && len(leftside) != 0 {
		leftside = sort(leftside)
	}
	if len(rightside) != 1 && len(rightside) != 0 {
		rightside = sort(rightside)
	}
	result := append(leftside, middle...)
	result = append(result, rightside...)
	return result
}
