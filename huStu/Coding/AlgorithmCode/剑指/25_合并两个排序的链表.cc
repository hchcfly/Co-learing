class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        ListNode* newHead = new ListNode(-1);
        ListNode* now = newHead;
        ListNode* temp = nullptr;
        while(list1!=nullptr && list2!=nullptr)
        {
            //选择一个要连接的节点
            if(list1->val <= list2->val)
            {
                temp = list1;
                list1 = list1->next;
            }
            else
            {
                temp = list2;
                list2 = list2->next;
            }
            now->next = temp;
            now = now->next;
        }
        if(list1 == nullptr)
        {
            now->next = list2;
        }
        if(list2 == nullptr)
        {
            now->next = list1;
        }
        return newHead->next;
    }
};