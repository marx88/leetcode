package lc00003

func lengthOfLongestSubstring(s string) int {
	m := 0

	// si:子串在字符串中的索引 ei:当前字符在子串中的索引
	for i, si := 0, 0; i < len(s); i++ {
		ei := strIndex(s[si:i], s[i:i+1])
		if ei > -1 {
			si = si + ei + 1 // 重复则重新指定si
		} else if m < i-si+1 {
			m = i - si + 1 // 新的最大长度已出现
		}
	}

	return m
}

// strIndex 字符在字符串中的索引
func strIndex(str, subStr string) int {
	for i, length := 0, len(str); i < length; i++ {
		if str[i:i+1] == subStr {
			return i
		}
	}
	return -1
}
