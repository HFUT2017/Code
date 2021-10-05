class MinStack {
    stack<int>stack1,stack2;
public:
    /** initialize your data structure here. */
    MinStack() {
        while(!stack1.empty()){
            stack1.pop();
        }

        while(!stack2.empty()){
            stack2.pop();
        }
    }
    
    void push(int x) {
        stack2.push(x);
        if(stack1.empty()||x<=stack1.top()){
            stack1.push(x);
        }
    }
    
    void pop() {
        if(!stack2.empty()){
            if(stack2.top()==stack1.top()){
                stack1.pop();
                stack2.pop();
            }else{
                stack2.pop();
            }
        }
    }
    
    int top() {
        return stack2.top();
    }
    
    int min() {
        return stack1.top();
    }
};
