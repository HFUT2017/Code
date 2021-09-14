class Solution {
public:
    ListNode* deleteNode(ListNode* head, int val) {
        ListNode *pre=new ListNode(-1);
        pre->next=head;
        while(pre->next->val!=val){
            pre=pre->next;
        }
        if(pre->next==head){
            return pre->next->next;
        }else{
            pre->next=pre->next->next;
        }
        return head;
    }
};
