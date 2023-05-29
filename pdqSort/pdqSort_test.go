package pdqSort

import (
	"math/rand"
	"testing"
	"time"
)

var (
	testSize int
	nums     []int
)

func initNums() {
	testSize = 64
	nums = make([]int, testSize)
	rand.Seed(time.Now().Unix())
}

func randNums() {
	for i := 0; i < testSize; i++ {
		nums[i] = rand.Int()
	}
}

func sortedNums() {
	for i := 0; i < testSize; i++ {
		nums[i] = i + 1
	}
}

func BenchmarkPdqSortRandom(b *testing.B) {
	initNums()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randNums()
		PdqSort(nums, 0, testSize-1)
	}
}

func BenchmarkPdqSortSorted(b *testing.B) {
	initNums()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortedNums()
		PdqSort(nums, 0, testSize-1)
	}
}
