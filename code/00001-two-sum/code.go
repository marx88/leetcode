package main

import (
	"log"
)

func main() {
	log.Println(twoSum([]int{2, 7, 11, 15}, 9))
	log.Println(twoSumMap([]int{2, 7, 11, 15}, 9))
}

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
