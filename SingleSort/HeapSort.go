package SingleSort

func HeapSort(nums []int) {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		shiftDown(nums, i, n)
	}
	for i := 0; i < n; i++ {
		swap(&nums[0], &nums[n-1-i])
		shiftDown(nums, 0, n-1-i)
	}
}

func shiftDown(nums []int, index int, n int) {
	for true {
		leftChild := index*2 + 1
		rightChild := index*2 + 2
		maxIndex := index
		if leftChild < n && nums[leftChild] > nums[maxIndex] {
			maxIndex = leftChild
		}
		if rightChild < n && nums[rightChild] > nums[maxIndex] {
			maxIndex = rightChild
		}
		if maxIndex == index {
			break
		}
		swap(&nums[maxIndex], &nums[index])
		index = maxIndex
	}
}
