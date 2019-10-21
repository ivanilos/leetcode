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
    ListNode* deleteDuplicates(ListNode* head) {
        ListNode* new_head = nullptr;
        ListNode* new_tail = nullptr;
        
        while(head != nullptr) {
            if (head->next != nullptr && head->val == head->next->val) {
                int val = head->val;
                while(head != nullptr && head->val == val) {
                    ListNode* aux = head;
                    head = head->next;
                    //free(aux);
                }
            } else {
                if (new_head == nullptr) {
                    new_head = head;
                } else {
                    new_tail->next = head;
                }
                new_tail = head;
                head = head->next;
                new_tail->next = nullptr;
            }
        }
        return new_head;
    }
};