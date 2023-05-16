package array

func twoSum(nums []int, target int) []int {
	seen := map[int]int{}
	for i, n := range nums {
		if _, ok := seen[target-n]; ok {
			return []int{i, seen[target-n]}
		}
		seen[n] = i
	}
	return []int{}
}
