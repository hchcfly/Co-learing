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
    ListNode* deleteNode(ListNode* head, int val) 
    {
        ListNode *rh = new ListNode(-1);
        rh->next = head;
        ListNode *pre = rh;
        ListNode *now = head;

        while(now != nullptr)
        {
            if(now->val == val)
            {
                pre->next = now->next;
                return rh->next;
            }
            pre = now;
            now = now->next;
        }
        return rh->next;
    }
};