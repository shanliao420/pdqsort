package SingleSort

func mergeNums(nums []int, start, mid, end int) {
	lStart := start
	rStart := mid + 1
	sortedNum := make([]int, end-start+1)
	mergePointer := 0
	for lStart <= mid && rStart <= end {
		if nums[lStart] <= nums[rStart] {
			sortedNum[mergePointer] = nums[lStart]
			mergePointer++
			lStart++
		} else {
			sortedNum[mergePointer] = nums[rStart]
			mergePointer++
			rStart++
		}
	}
	for lStart <= mid {
		sortedNum[mergePointer] = nums[lStart]
		mergePointer++
		lStart++
	}

	for rStart <= end {
		sortedNum[mergePointer] = nums[rStart]
		mergePointer++
		rStart++
	}
	for i := start; i <= end; i++ {
		nums[i] = sortedNum[i-start]
	}
}

func MergeSort(nums []int, start int, end int) {
	if start == end {
		return
	}

	mid := (start + end) / 2
	MergeSort(nums, start, mid)
	MergeSort(nums, mid+1, end)

	mergeNums(nums, start, mid, end)
}
