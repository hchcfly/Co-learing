class CQueue {
private:
    stack<int> s1;
    stack<int> s2;
public:
    CQueue() {
        while(!s1.empty())
        {
            s1.pop();
        }
        while(!s2.empty())
        {
            s2.pop();
        }
    }
    
    void appendTail(int value) 
    {
        s1.push(value);
    }
    
    int deleteHead() 
    {
        if(s2.empty())
        {
            if(s1.empty())
            {
                return -1;
            }
            else
            {
                while(!s1.empty())
                {
                    s2.push(s1.top());
                    s1.pop();
                }
            }
        }
        int headvalue = s2.top();
        s2.pop();
        return headvalue;
    }
};

/**
 * Your CQueue object will be instantiated and called as such:
 * CQueue* obj = new CQueue();
 * obj->appendTail(value);
 * int param_2 = obj->deleteHead();
 */