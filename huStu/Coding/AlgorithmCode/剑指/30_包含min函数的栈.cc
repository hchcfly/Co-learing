class MinStack {
private:
    stack<int> s1;
    stack<int> s_min;
public:
    /** initialize your data structure here. */
    MinStack() {
        while(!s1.empty())
        {
            s1.pop();
        }
        while(!s_min.empty())
        {
            s_min.pop();
        }
    }
    
    void push(int x) 
    {
        s1.push(x);
        if(s_min.empty())
        {
            s_min.push(x);
        }   
        else
        {
            s_min.push(s_min.top()>x?x:s_min.top());
        }
    }
    
    void pop() {
        s1.pop();
        s_min.pop();
    }
    
    int top() {
        return s1.top();
    }
    
    int min() {
        return s_min.top();
    }
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack* obj = new MinStack();
 * obj->push(x);
 * obj->pop();
 * int param_3 = obj->top();
 * int param_4 = obj->min();
 */