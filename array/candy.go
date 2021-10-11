package array

// candy 糖果分配, 每个孩子至少有一颗
// Input: [1,0,2]
// Output: 5
func candy(ratings []int) []int {
	if len(ratings) == 0 {
		return nil
	}
	// 每个孩子至少有一颗
	nums := make([]int, len(ratings))
	for i := 0; i < len(nums); i++ {
		nums[i] = 1
	}
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] > ratings[i+1] {
			nums[i] += 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] > ratings[i-1] {
			nums[i] += 1
		}
	}
	return nums
}
