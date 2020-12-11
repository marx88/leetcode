<?php

namespace leetcode\lc00001;

class TwoSum
{
    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    public function twoSum($nums, $target)
    {
        $len = count($nums);
        $map = [];
        for ($i = 0; $i < $len; $i++) {
            $key = $target - $nums[$i];
            if (array_key_exists($key, $map)) {
                return [$map[$key], $i];
            }
            $map[$nums[$i]] = $i;
        }
        return [];
    }
}
