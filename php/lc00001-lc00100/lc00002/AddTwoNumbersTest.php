<?php

namespace leetcode\lc00002;

use PHPUnit\Framework\TestCase;

class AddTwoNumbersTest extends TestCase
{
    public function testAddTwoNumbers()
    {
        $cases = [
            ['l1' => $this->ln([9, 4, 2]), 'l2' => $this->ln([9, 4, 6, 5]), 'expected' => 10407],
            ['l1' => $this->ln([1, 0]), 'l2' => $this->ln([1, 0, 2]), 'expected' => 112],
            ['l1' => $this->ln([1]), 'l2' => $this->ln([0]), 'expected' => 1],
            ['l1' => $this->ln([0]), 'l2' => $this->ln([0]), 'expected' => 0],
        ];

        $obj = new AddTwoNumbers();
        for ($i = 0; $i < count($cases); $i++) {
            $this->assertEquals($cases[$i]['expected'], $this->ln2Int($obj->addTwoNumbers($cases[$i]['l1'], $cases[$i]['l2'])));
        }
    }

    public function ln($nums)
    {
        $ln = null;
        for ($i = 0; $i < count($nums); $i++) {
            $ln = new ListNode($nums[$i], $ln);
        }
        return $ln;
    }

    public function ln2Int($ln)
    {
        $num = 0;
        $i = 1;
        while ($ln !== null) {
            $num += ($ln->val * $i);
            $ln = $ln->next;
            $i *= 10;
        }
        return $num;
    }
}
