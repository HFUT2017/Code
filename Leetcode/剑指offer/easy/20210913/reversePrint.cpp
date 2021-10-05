class Solution {
public:
    vector<int> reversePrint(ListNode* head) {
        vector<int> res;
        while(head!=NULL){
            res.push_back(head->val);
            head=head->next;
        }
        for(int i=0;i<res.size()/2;i++){
            swap(res[i],res[res.size()-i-1]);
        }
        return res;
    }
};
