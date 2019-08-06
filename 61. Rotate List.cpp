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
    ListNode* rotateRight(ListNode* head, int k) {
        if (head == nullptr) return nullptr;
        int len = get_len(head);
        k %= len;
        
        ListNode* new_head = head;
        for (int i = 0; i < len - k - 1; i++) {
            new_head = new_head->next;
        }
        
        ListNode* one_ahead = new_head->next;
        new_head->next = nullptr;
        new_head = one_ahead;
        if (new_head == nullptr) { new_head = head; }
        
        while(one_ahead != nullptr) {
            if (one_ahead->next == nullptr) {
                one_ahead->next = head;
                break;
            }
            one_ahead = one_ahead->next;
        }
        return new_head;
    }
    
private:
    int get_len(ListNode* start) {
        int len = 0;
        while(start != nullptr) {
            len++;
            start = start->next;
        }
        return len;
    }
};