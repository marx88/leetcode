<?php
class Solution
{
    /**
     * @param String $s
     * @return Integer
     */
    public function lengthOfLongestSubstring($s)
    {
        $m = 0;
        for ($i = 0, $l = strlen($s), $si = 0; $i < $l; $i++) {
            $ei = strpos(substr($s, $si, $i - $si), substr($s, $i, 1));
            if ($ei !== false) {
                $si = $si + $ei + 1;
            } elseif ($m < $i - $si + 1) {
                $m = $i - $si + 1;
            }
        }
        return $m;
    }
}

echo (new Solution())->lengthOfLongestSubstring('pwwkew');
