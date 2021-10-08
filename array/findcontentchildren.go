package array

import "sort"

// findContentChildren 饼干分配
// Input: [1,2], [1,2,3]
// Output: 2
func findContentChildren(children, cookies []int) int {
	if len(children) == 0 || len(cookies) == 0 {
		return 0
	}
	child, cookie := 0, 0
	sort.Ints(children)
	sort.Ints(cookies)
	for child < len(children) && cookie < len(cookies) {
		if children[child] <= cookies[cookie] {
			child++
		}
		cookie++
	}
	return child
}
