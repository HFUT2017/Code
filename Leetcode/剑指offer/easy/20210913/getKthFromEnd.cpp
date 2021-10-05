class Solution {
public:
    ListNode* getKthFromEnd(ListNode* head, int k) {
        ListNode *pre=new ListNode(-1);
        pre->next=head;
        while(k--){
            if(pre==NULL){
                return NULL;
            }
            pre=pre->next;
        }
        while(pre->next!=NULL){
            head=head->next;
            pre=pre->next;
        }
        return head;
    }
};
