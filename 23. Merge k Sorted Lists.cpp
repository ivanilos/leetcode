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
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        ListNode* head = nullptr;
        ListNode* tail = nullptr;
        
        priority_queue<pair<int, ListNode*>> pq;
        for (auto it : lists) {
            if (it != nullptr) {
                pq.push({-it->val, it});
            }
        }
        
        while(!pq.empty()) {
            ListNode* nextNode = pq.top().second;
            pq.pop();
            
            if (head == nullptr) {
                head = nextNode;
                tail = head;
            } else {
                tail->next = nextNode;
                tail = tail->next;
            }
            if (nextNode->next != nullptr) {
                pq.push({-(nextNode->next)->val, nextNode->next});
            }
            tail->next = nullptr;
        }
        return head;
    }
};