/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* next;
    Node* random;
    
    Node(int _val) {
        val = _val;
        next = NULL;
        random = NULL;
    }
};
*/
class Solution {
public:
    Node* copyRandomList(Node* head) {
    if(head == nullptr) return nullptr;
    //复制链表
    Node* lhead = head;
    while(lhead!=nullptr)
    {
        Node* addnode = new Node(lhead->val);
        addnode->next = lhead->next;
        lhead->next = addnode;
        lhead = lhead->next->next;
    }
    
    //对random进行处理
    Node* prehead = head;
    while(prehead!=nullptr && prehead->next!=nullptr)
    {
        Node* randnode = prehead->random;
        if(randnode != nullptr) prehead->next->random = randnode->next;
        else prehead->next->random = nullptr;
        prehead = prehead->next->next;
    }
    // Node* lookhead = head;
    // while(lookhead!=nullptr)
    // {
    //     if(lookhead->random != nullptr)
    //         cout << lookhead->val << " : " << lookhead->random->val << endl;
    //     lookhead = lookhead->next;
    // }

    //拆分链表
    Node* headl = head;
    Node* headr = head->next;
    Node* ansnode = headr;
    cout << ansnode->val << endl;
    while(headl!=nullptr && headr!=nullptr && headr->next != nullptr)
    {
        headl->next = headr->next;
        headl = headl->next;
        headr->next = headl->next;
        headr = headr->next;
    }
    //注意点:对与原链表的处理,否则报错
    headl->next = nullptr;  
    return ansnode;
    }
};