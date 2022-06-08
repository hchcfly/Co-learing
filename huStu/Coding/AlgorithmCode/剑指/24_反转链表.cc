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
        ListNode* cur = head;
        ListNode* pre  = nullptr;
        ListNode* nnode = nullptr;
        while(cur!=nullptr) 
        {
            nnode = cur->next;
            cur->next = pre;
            pre = cur;
            cur = nnode;
        }
        return pre;
    }
};