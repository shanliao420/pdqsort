package SingleSort

import (
	"math/rand"
	"testing"
	"time"
)

var (
	nums     []int
	testSize int
)

func TestHeapSort(t *testing.T) {
	HeapSort(nums)
}

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

func BenchmarkHeapSortRandom(b *testing.B) {
	initNums()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randNums()
		HeapSort(nums)
	}
}

func BenchmarkQuickSortRandom(b *testing.B) {
	initNums()
	n := len(nums) - 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randNums()
		QuickSort(nums, 0, n)
	}
}

func BenchmarkInsertionSortRandom(b *testing.B) {
	initNums()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randNums()
		InsertionSort(nums)
	}
}

func BenchmarkHeapSortSorted(b *testing.B) {
	initNums()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortedNums()
		HeapSort(nums)
	}
}

func BenchmarkQuickSortSorted(b *testing.B) {
	initNums()
	n := len(nums) - 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortedNums()
		QuickSort(nums, 0, n)
	}
}

func BenchmarkInsertionSortSorted(b *testing.B) {
	initNums()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortedNums()
		InsertionSort(nums)
	}
}
