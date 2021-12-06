package day1

func one(nums []int64) int {
	numLarger := 0
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		if curr > prev {
			numLarger += 1
		}
		prev = curr
	}
	return numLarger
}

// Consider sums of a three-measurement sliding window. How many sums are larger than the previous sum?
func two(nums []int64) int {
	numLarger := 0
	for i := 0; i + 3 < len(nums); i++ {
		firstWindow := nums[i] + nums[i+1] + nums[i+2]
		secondWindow := nums[i+1] + nums[i+2] + nums[i+3]
		if secondWindow > firstWindow {
			numLarger += 1
		}
	}
	return numLarger
}


