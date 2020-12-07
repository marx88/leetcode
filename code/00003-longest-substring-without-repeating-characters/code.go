package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(lengthOfLongestSubstring3("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	var t map[string]bool
	l := len(s)
	m := 0
	for i, j, k := 0, 0, ""; i < l; i++ {
		t = map[string]bool{s[i : i+1]: true}
		for j = i + 1; j < l; j++ {
			k = s[j : j+1]
			if _, ok := t[k]; ok {
				break
			}
			t[k] = true
		}
		if m < j-i {
			m = j - i
		}
		if j+1 == l {
			break
		}
	}
	return m
}

func lengthOfLongestSubstring2(s string) int {
	m, l, t := 0, len(s), ""
	for i := 0; i < l; i++ {
		b := s[i : i+1]
		if si := strings.Index(t, b); si > -1 {
			if m < len(t) {
				m = len(t)
			}
			t = t[si+1:] + b
		} else {
			t += b
		}
	}
	if len(t) > m {
		m = len(t)
	}
	return m
}

func lengthOfLongestSubstring3(s string) int {
	m := 0

	// si:子串在字符串中的索引 ei:当前字符在子串中的索引
	for i, si := 0, 0; i < len(s); i++ {
		ei := strings.Index(s[si:i], s[i:i+1])
		if ei > -1 {
			si = si + ei + 1 // 重复则重新指定si
		} else if m < i-si+1 {
			m = i - si + 1 // 新的最大长度已出现
		}
	}

	return m
}
