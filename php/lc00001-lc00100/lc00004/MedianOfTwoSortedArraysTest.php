<?php

namespace leetcode\lc00004;

use PHPUnit\Framework\TestCase;

class MedianOfTwoSortedArraysTest extends TestCase
{
    public function testFindMedianSortedArrays()
    {
        $cases = [
            ["nums1" => [1, 3], "nums2" => [2], "expected" => 2],
            ["nums1" => [1, 2], "nums2" => [3, 4], "expected" => 2.5],
            ["nums1" => [0, 0], "nums2" => [0, 0], "expected" => 0],
            ["nums1" => [], "nums2" => [1], "expected" => 1],
            ["nums1" => [2], "nums2" => [], "expected" => 2],
        ];

        $obj = new MedianOfTwoSortedArrays();
        for ($i = 0; $i < count($cases); $i++) {
            $this->assertEquals($cases[$i]['expected'], $obj->findMedianSortedArrays($cases[$i]['nums1'], $cases[$i]['nums2']));
        }
    }
}
