class Solution {
public:
    int res,t;
    void preorder(TreeNode *root){
        if(root==NULL||t==0){
            return;
        }
        preorder(root->right);
        if(--t==0){
            res=root->val;
            return;
        }
        preorder(root->left);
    }
    int kthLargest(TreeNode* root, int k) {
        t=k;
        preorder(root);
        return res;
    }
};
