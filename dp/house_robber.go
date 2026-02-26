package dp

func rob(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		if b > a {
			return b

		}
		return a
	}
	computed_results := map[int]int{}

	// start robbing
	var max_robbed func(i int) int
	max_robbed = func(i int) int {
		if i == 0 {
			return nums[0]
		}
		if i < 0 {
			return 0
		}

		if v, ok := computed_results[i]; ok {
			return v

		}
		computed_results[i] = max(max_robbed(i-2)+nums[i], max_robbed(i-1))
		return computed_results[i]
	}
	return max_robbed(len(nums) - 1)

}
