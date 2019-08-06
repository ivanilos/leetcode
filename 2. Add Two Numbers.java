/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode ans = null;
        ListNode lastDig = null;
        int carry = 0;
        while(l1 != null || l2 != null || carry != 0) {
            int dig1 = 0;
            int dig2 = 0;
            if (l1 != null) {
                dig1 = l1.val;
                l1 = l1.next;
            }
            if (l2 != null) {
                dig2 = l2.val;
                l2 = l2.next;
            }
            
            ListNode nextDig = new ListNode((dig1 + dig2 + carry) % 10);
            nextDig.next = null;
            if (ans == null) {
                ans = nextDig;
                lastDig = nextDig;
            } else {
                lastDig.next = nextDig;
                lastDig = nextDig;
            }   
            carry = (dig1 + dig2 + carry) / 10;
        }
        return ans;
    }
}