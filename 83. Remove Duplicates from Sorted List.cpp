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
        if (head == nullptr) {
            return nullptr;
        }
        ListNode* currentNode = head;
        ListNode* nextNode = head->next;
        while(nextNode != nullptr) {
            if (currentNode->val == nextNode->val) {
                currentNode->next = nextNode->next;
                free(nextNode);
                nextNode = currentNode->next;
            } else {
                currentNode = nextNode;
                nextNode = nextNode->next;
            }
        }
        return head;
    }
};