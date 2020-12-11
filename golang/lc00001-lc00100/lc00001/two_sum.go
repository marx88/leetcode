package lc00001

func twoSum(nums []int, target int) []int {
	for i, lenNums := 0, len(nums); i < lenNums; i++ {
		for j := i + 1; j < lenNums; j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}

	return nil
}

// twoSumMap map
func twoSumMap(nums []int, target int) []int {
	var key int
	m := make(map[int]int, 0)
	for i, v := range nums {
		key = target - v
		if j, ok := m[key]; ok {
			return []int{j, i}
		}
		m[v] = i
	}

	return nil
}
