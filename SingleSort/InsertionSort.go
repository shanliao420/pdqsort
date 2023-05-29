package SingleSort

func InsertionSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
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
