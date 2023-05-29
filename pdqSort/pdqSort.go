package pdqSort

var (
	limit = 10
)

func PdqSort(nums []int, left int, right int) {
	if (right - left + 1) <= 24 {
		insertionSort(nums, left, right)
		return
	}
	if limit <= 0 {
		heapSort(nums, left, right)
		return
	}
	quickSort(nums, left, right)
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func getMid(nums []int, left int, right int) int {
	mid := (left + right) / 2
	condition := mid
	if nums[left] > nums[mid] && nums[right] > nums[left] {
		mid = left
	}
	if nums[right] > nums[mid] && nums[left] > nums[right] {
		mid = right
	}
	if condition == mid {
		return -1
	}
	return mid
}

func partition(nums []int, left int, right int) int {
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

func insertionSort(nums []int, start int, end int) {
	for i := start + 1; i < end; i++ {
		x := nums[i]
		j := i - 1
		for j >= 0 {
			if nums[j] >= x {
				nums[j+1] = nums[j]
				j--
			} else {
				break
			}
		}
		nums[j+1] = x
	}
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	mid := getMid(nums, left, right)
	if mid == -1 {
		insertionSort(nums, left, right)
		return
	}
	swap(&nums[left], &nums[mid])
	pivot := partition(nums, left, right)
	if pivot < (left+right)/8 {
		limit--
	}
	PdqSort(nums, left, pivot-1)
	PdqSort(nums, pivot+1, right)
}

func heapSort(nums []int, start int, end int) {
	n := end + 1
	for i := n - 1; i >= start; i-- {
		shiftDown(nums, i, n)
	}
	for i := start; i < n; i++ {
		swap(&nums[start], &nums[n-1-i])
		shiftDown(nums, start, n-1-i)
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
