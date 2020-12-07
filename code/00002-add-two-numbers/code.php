<?php
class ListNode
{
    public $val = 0;
    public $next = null;
    public function __construct($val = 0, $next = null)
    {
        $this->val = $val;
        $this->next = $next;
    }
}

class Solution
{

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    public function addTwoNumbers($l1, $l2)
    {
        $first = new ListNode($l1->val + $l2->val, null);
        $temp = $first;
        while ($l1->next || $l2->next || $temp->val > 9) {
            $temp->next = new ListNode(intval($temp->val / 10), null);
            $temp->val = $temp->val % 10;
            $temp = $temp->next;
            if ($l1->next !== null) {
                $temp->val += $l1->next->val;
                $l1 = $l1->next;
            }
            if ($l2->next !== null) {
                $temp->val += $l2->next->val;
                $l2 = $l2->next;
            }
        }
        return $first;
    }
}

$l1 = new ListNode(2, new ListNode(4, new ListNode(9)));
$l2 = new ListNode(5, new ListNode(6, new ListNode(4, new ListNode(9))));
print_r((new Solution())->addTwoNumbers($l1, $l2));
