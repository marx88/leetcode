<?php

namespace leetcode\lc00001;

use PHPUnit\Framework\TestCase;

class TwoSumTest extends TestCase
{
    public function testTwoSum()
    {
        $cases = [
            ['nums' => [2, 7, 11, 15], 'target' => 9, 'expected' => [0, 1]],
            ['nums' => [2, 7, 11, 15], 'target' => 17, 'expected' => [0, 3]],
        ];

        $obj = new TwoSum();
        for ($i = 0; $i < count($cases); $i++) {
            $this->assertEquals($cases[$i]['expected'], $obj->twoSum($cases[$i]['nums'], $cases[$i]['target']));
        }
    }
}
