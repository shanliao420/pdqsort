package SingleSort

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func swapMid(nums []int, left int, right int) {
	mid := (left + right) / 2
	if nums[left] > nums[mid] && nums[right] > nums[left] {
		mid = left
	}
	if nums[right] > nums[mid] && nums[left] > nums[right] {
		mid = right
	}
	swap(&nums[left], &nums[mid])
}

func partition(nums []int, left int, right int) int {
	swapMid(nums, left, right)
	pivot := left
	base := nums[pivot]
	for left < right {
		for left < right && nums[right] >= base {
			right--
		}
		for left < right && nums[left] <= base {
			left++
		}
		swap(&nums[left], &nums[right])
	}
	swap(&nums[left], &nums[pivot])
	return left
}

//func partition(nums []int, left int, right int) int {
//	swapMid(nums, left, right)
//	pivot := left
//	minThan := left + 1
//	for i := left + 1; i <= right; i++ {
//		if nums[pivot] >= nums[i] {
//			swap(&nums[minThan], &nums[i])
//			minThan++
//		}
//	}
//	swap(&nums[pivot], &nums[minThan-1])
//	pivot = minThan - 1
//	return pivot
//}

func QuickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	QuickSort(nums, left, pivot-1)
	QuickSort(nums, pivot+1, right)
}
