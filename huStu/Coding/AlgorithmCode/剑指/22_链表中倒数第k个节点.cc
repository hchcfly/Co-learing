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
    ListNode* getKthFromEnd(ListNode* head, int k) 
    {
        ListNode *pre = head;
        ListNode *now = head;
        int count = k-1;
        while (count--)
        {
            now = now->next;
        }
        while(now->next != nullptr)
        {
            pre = pre->next;
            now = now->next;
        }
        return pre;
    }
};