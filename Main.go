package main

import (
	"fmt"
	"math/rand"
	"time"
	"work.tangthinker/sort_algorithm/pdqSort"
)

func main() {
	testSize := 64
	nums := make([]int, testSize)
	rand.Seed(time.Now().Unix())
	for i := 0; i < testSize; i++ {
		nums[i] = rand.Int() % 100
	}
	fmt.Println(nums)
	pdqSort.PdqSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
