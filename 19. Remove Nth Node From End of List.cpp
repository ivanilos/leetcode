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
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        if (head == nullptr) {
            return nullptr;
        }
        int sz = getListSize(head);
        return removeNthFromBeginning(head, sz - n);
    }
    
    ListNode* removeNthFromBeginning(ListNode* head, int n) {
        if (head == nullptr) {
            return nullptr;
        }
        if (n == 0) {
            ListNode* toRemove = head;
            head = head->next;
            free(toRemove);
            return head;
        } else {
            ListNode* newHead = head;
            while(n > 1) {
                head = head->next;
                n--;
            }
            ListNode* toRemove = head->next;
            head->next = toRemove->next;
            free(toRemove);
            return newHead;
        }
    }
    
    int getListSize(ListNode* head) {
        int sz = 0;
        while(head != nullptr) {
            sz++;
            head = head->next;
        }
        return sz;
    }
};