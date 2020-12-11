package lc00004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1, length2 := len(nums1), len(nums2)
	for index1, index2, offset := length1/2, (length1+length2)/2-length1/2, (min(length1, length2)/4)+1; ; offset = (offset + 1) / 2 {
		if (index1-1 >= 0 && index1-1 < length1) && (index2 >= 0 && index2 < length2) && nums1[index1-1] > nums2[index2] { // 数组1索引如果还能且需要左移
			index1 -= offset
			index2 += offset
		} else if (index1 >= 0 && index1 < length1) && (index2-1 >= 0 && index2-1 < length2) && nums1[index1] < nums2[index2-1] { // 数组1索引如果还能且需要右移
			index1 += offset
			index2 -= offset
		} else if (length1+length2)%2 == 0 { // 偶数个 求[左面大的 右面小的]均值
			return (getValue(nums1, index1-1, nums2, index2-1, true) + getValue(nums1, index1, nums2, index2, false)) / 2
		} else { // 奇数个 取右面小的
			return getValue(nums1, index1, nums2, index2, false)
		}
	}
}

// getValue 根据两个数组及给定的索引 求存在的最大/最小值
func getValue(nums1 []int, index1 int, nums2 []int, index2 int, bigger bool) (num float64) {
	isSet := false
	if index1 >= 0 && index1 < len(nums1) {
		num = float64(nums1[index1])
		isSet = true
	}
	if index2 >= 0 && index2 < len(nums2) && (!isSet || (bigger && num < float64(nums2[index2])) || (!bigger && num > float64(nums2[index2]))) {
		num = float64(nums2[index2])
	}
	return
}

// min 获取短数组的长度 golang没有三目运算 只能这样
//
// leetcode上有人嵌套调用解决获取短数组长度的问题 即判断长度 如果第一个是长的 把两个数组调换位置再传参调用即可
func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}
