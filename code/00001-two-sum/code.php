<?php

/**
 * @param Integer[] $nums
 * @param Integer $target
 * @return Integer[]
 */
function twoSum($nums, $target)
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

print_r(twoSum([2, 7, 11, 15], 9));
