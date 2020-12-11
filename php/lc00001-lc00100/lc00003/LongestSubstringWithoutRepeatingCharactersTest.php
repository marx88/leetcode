<?php

namespace leetcode\lc00003;

use PHPUnit\Framework\TestCase;

class LongestSubstringWithoutRepeatingCharactersTest extends TestCase
{
    public function testLengthOfLongestSubstring()
    {
        $cases = [
            ["str" => "abcabcbb", "expected" => 3],
            ["str" => "bbbbb", "expected" => 1],
            ["str" => "pwwkew", "expected" => 3],
            ["str" => "", "expected" => 0],
        ];

        $obj = new LongestSubstringWithoutRepeatingCharacters();
        for ($i = 0; $i < count($cases); $i++) {
            $this->assertEquals($cases[$i]['expected'], $obj->lengthOfLongestSubstring($cases[$i]['str']));
        }
    }
}
