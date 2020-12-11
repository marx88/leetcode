<?php

namespace leetcode\lc00004;

class MedianOfTwoSortedArrays
{
    /**
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @return Float
     */
    public function findMedianSortedArrays($nums1, $nums2)
    {
        $length1 = count($nums1);
        $length2 = count($nums2);
        $index1 = intval($length1 / 2);
        $index2 = intval(($length1 + $length2) / 2) - intval($length1 / 2);
        $offset = intval($this->min($length1, $length2) / 4) + 1;
        while (true) {
            if (($index1 - 1 >= 0 && $index1 - 1 < $length1) && ($index2 >= 0 && $index2 < $length2) && $nums1[$index1 - 1] > $nums2[$index2]) { // 数组1索引如果还能且需要左移
                $index1 -= $offset;
                $index2 += $offset;
            } else if (($index1 >= 0 && $index1 < $length1) && ($index2 - 1 >= 0 && $index2 - 1 < $length2) && $nums1[$index1] < $nums2[$index2 - 1]) { // 数组1索引如果还能且需要右移
                $index1 += $offset;
                $index2 -= $offset;
            } else if (($length1 + $length2) % 2 == 0) { // 偶数个 求[左面大的 右面小的]均值
                return ($this->getValue($nums1, $index1 - 1, $nums2, $index2 - 1, true) + $this->getValue($nums1, $index1, $nums2, $index2, false)) / 2;
            } else { // 奇数个 取右面小的
                return $this->getValue($nums1, $index1, $nums2, $index2, false);
            }
            $offset = intval(($offset + 1) / 2);
        }
    }

    // getValue 根据两个数组及给定的索引 求存在的最大/最小值
    private function getValue($nums1, $index1, $nums2, $index2, $bigger)
    {
        $isSet = false;
        $num = 0;
        if ($index1 >= 0 && $index1 < count($nums1)) {
            $num = $nums1[$index1];
            $isSet = true;
        }
        if ($index2 >= 0 && $index2 < count($nums2) && (!$isSet || ($bigger && $num < $nums2[$index2]) || (!$bigger && $num > $nums2[$index2]))) {
            $num = $nums2[$index2];
        }
        return $num;
    }

    // min 获取短数组的长度 golang没有三目运算 只能这样
    // leetcode上有人嵌套调用解决获取短数组长度的问题 即判断长度 如果第一个是长的 把两个数组调换位置再传参调用即可
    private function min($num1, $num2)
    {
        if ($num1 > $num2) {
            return $num2;
        }
        return $num1;
    }
}
