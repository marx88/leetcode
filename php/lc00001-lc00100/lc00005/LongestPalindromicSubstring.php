<?php

namespace leetcode\lc00005;

class LongestPalindromicSubstring
{
    /**
     * @param String $s
     * @return String
     */
    public function longestPalindrome($s)
    {
        $si = $maxLen = 0;
        for ($l = $r = $m = 0, $len = strlen($s); $m < $len;) {
            // 下面这个循环解决奇偶回文长度问题 同时优化了N个重复字符问题
            for ($l = $m - 1, $r = $m; $r < $len && $s[$m] == $s[$r]; $r++) {
            }

            // 中心扩散
            for ($m = $r; $l >= 0 && $r < $len && $s[$l] == $s[$r]; $l--, $r++) {
            }

            // 判断是否新的最大回文
            if ($maxLen < $r - $l - 1) {
                $si = $l + 1;
                $maxLen = $r - $l - 1;
            }
        }
        return substr($s, $si, $maxLen);
    }

    /**
     * Manacher[马拉车]算法
     * @param String $s
     * @return String
     */
    public function manacher($s)
    {
        $si = $maxLen = 0;
        for ($p = [], $len = strlen($s) * 2 + 1, $i = $l = $r = $maxRIdx = $cIdx = 0; $i < $len; $i++) {
            if ($i < $maxRIdx) {
                $p[$i] = $p[2 * $cIdx - $i];
                if ($p[$i] > $maxRIdx - $i) {
                    $p[$i] = $maxRIdx - $i;
                    continue;
                } else if ($p[$i] < $maxRIdx - $i) {
                    continue;
                }
                $l = $i - ($p[$i] + 1);
                $r = $i + ($p[$i] + 1);
            } else {
                $p[$i] = 0;
                $l = $i - 1;
                $r = $i + 1;
            }

            for (; $l >= 0 && $r < $len && ($r % 2 == 0 || $s[intval($l / 2)] == $s[intval($r / 2)]); $l--, $r++, $p[$i]++) {}

            if ($i + $p[$i] > $maxRIdx) {
                $cIdx = $i;
                $maxRIdx = $i + $p[$i];
            }

            if ($p[$i] > $maxLen) {
                $si = intval(($l + 1) / 2);
                $maxLen = $p[$i];
            }
        }
        return substr($s, $si, $maxLen);
    }
}
