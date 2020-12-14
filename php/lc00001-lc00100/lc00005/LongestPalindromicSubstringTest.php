<?php

namespace leetcode\lc00005;

use PHPUnit\Framework\TestCase;

class LongestPalindromicSubstringTest extends TestCase
{
    private $cases = [
        ["s" => "babad", "expected" => "bab"],
        ["s" => "cbbd", "expected" => "bb"],
        ["s" => "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth", "expected" => "ranynar"],
        ["s" => "aacabdkacaa", "expected" => "aca"],
        ["s" => "edllaaaaallf", "expected" => "llaaaaall"],
    ];

    public function testLongestPalindrome()
    {
        $obj = new LongestPalindromicSubstring();
        for ($i = 0; $i < count($this->cases); $i++) {
            $this->assertEquals($this->cases[$i]['expected'], $obj->longestPalindrome($this->cases[$i]['s']));
        }
    }

    public function testManacher()
    {
        $obj = new LongestPalindromicSubstring();
        for ($i = 0; $i < count($this->cases); $i++) {
            $this->assertEquals($this->cases[$i]['expected'], $obj->manacher($this->cases[$i]['s']));
        }
    }
}
