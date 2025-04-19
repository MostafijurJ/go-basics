package problem_solving

func CountPairs(nums []int, k int) int {
	var ans = 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {

			var div = (i * j) % k

			if nums[i] == nums[j] && div == 0 {
				ans++
			}
		}
	}

	return ans
}
