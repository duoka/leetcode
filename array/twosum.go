package array

// TwoSum Given nums = [2, 7, 11, 15], target = 9
func TwoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	tmp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		key := target - nums[i]
		if _, ok := tmp[key]; ok {
			return []int{tmp[key], i}
		}
		tmp[nums[i]] = i
	}

	return nil
}
