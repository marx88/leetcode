package lc00005

// ==================== 暴力破解 ====================

func longestPalindrome(s string) string {
	sLen, reverseStr := len(s), reverseString(s)
	for i := sLen; i > 1; i-- {
		for j := 0; i+j <= sLen; j++ {
			if s[j:i+j] == reverseStr[sLen-j-i:sLen-j] {
				return s[j : i+j]
			}
		}
	}
	if sLen > 0 {
		return s[:1]
	}
	return ""
}

func reverseString(str string) string {
	bytes := []rune(str)
	for from, to := 0, len(bytes)-1; from < to; from, to = from+1, to-1 {
		bytes[from], bytes[to] = bytes[to], bytes[from]
	}
	return string(bytes)
}

// ==================== 中心扩散 ====================

func centerSpread(s string) string {
	var l, r, i, si, mLen int
	for sLen := len(s); i < sLen; {
		l, r = i-1, i

		// 下面这个循环解决奇偶回文长度问题 同时优化了N个重复字符问题
		for r < sLen && s[i] == s[r] {
			r++
		}
		i = r

		// 中心扩散
		for l >= 0 && r < sLen && s[l] == s[r] {
			l, r = l-1, r+1
		}

		// 判断是否新的最大回文
		if mLen < r-l-1 {
			mLen, si = r-l-1, l+1
		}
	}
	return s[si : si+mLen]
}

// ==================== Manacher[马拉车]算法 ====================

func manacher(s string) string {
	var sIdx, maxLen, lIdx, rIdx, cIdx, maxRIdx int
	for i, p, pLen := 0, make([]int, 2*len(s)+1), 2*len(s)+1; i < pLen; i++ {
		// 前方高能，非战斗人员请不要轻易进入该判断，更不要尝试解读里面的代码
		if i < maxRIdx {
			// 马拉车算法核心 概念:
			// 1.镜像长度[以镜像中心‘cIdx’为对称点，当前索引‘i’的对称索引里存储的回文长度]
			// 2.右侧长度[当前索引‘i’到最大右索引‘maxRIdx’的长度]
			p[i] = p[2*cIdx-i] // 获取镜像长度
			if p[i] > maxRIdx-i {
				// 镜像长度比右侧长度长，该索引的回文长度则等于右侧长度
				p[i] = maxRIdx - i
				continue
			} else if p[i] < maxRIdx-i {
				// 镜像长度比右侧长度短，该索引的回文长度则等于镜像长度
				continue
			}
			// 镜像长度等于右侧长度，该索引的回文长度则以镜像长度为基础，继续进行中心扩展
			lIdx, rIdx = i-(p[i]+1), i+(p[i]+1)
		} else {
			// 没有镜像长度，直接从当前索引两侧进行中心扩展
			lIdx, rIdx = i-1, i+1
		}

		// 中心扩展 如果右索引是2的整数倍则默认两个字符相等，否则需要判断原串中两个索引对应的字符是否相等
		for lIdx >= 0 && rIdx < pLen && (rIdx%2 == 0 || s[lIdx/2] == s[rIdx/2]) {
			p[i], lIdx, rIdx = p[i]+1, lIdx-1, rIdx+1
		}

		// 最大右索引、中心对称点需要更新
		if i+p[i] > maxRIdx {
			maxRIdx, cIdx = i+p[i], i
		}

		// 最大回文起始索引、最大回文长度需要更新
		if p[i] > maxLen {
			sIdx, maxLen = (lIdx+1)/2, p[i]
		}
	}
	return s[sIdx : sIdx+maxLen]
}
