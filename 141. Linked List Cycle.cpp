/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    bool hasCycle(ListNode *head) {
        ListNode* first = head;
        ListNode* sec = head;
        
        for (int max_len = 1;; max_len *= 2) {
            int walked = 0;
            while(sec != nullptr && walked < max_len) {
                sec = sec->next;
                walked++;
                if (first == sec) {
                    break;
                }
            }
            if (sec == nullptr) return false;
            if (first == sec) return true;

            first = sec;
        }
    }
};