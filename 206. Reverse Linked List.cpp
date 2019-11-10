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
    ListNode* reverseList(ListNode* head) {
        stack<ListNode*> st;
        
        while(head != nullptr) {
            st.push(head);
            head = head->next;
        }
        
        ListNode* new_head = nullptr;
        ListNode* new_tail = nullptr;
        while(!st.empty()) {
            ListNode* next_elem = st.top();
            st.pop();
            
            if (new_head == nullptr) {
                new_head = next_elem;
                new_tail = next_elem;
            } else {
                new_tail-> next = next_elem;
                new_tail = new_tail->next;
            }
            new_tail->next = nullptr;
        }
        return new_head;
    }
};