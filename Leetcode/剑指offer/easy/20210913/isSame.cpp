class Solution {
public:
    bool isSame(TreeNode* L,TreeNode* R){
        if(L==NULL&&R==NULL){
            return true;
        }
        if(L==NULL||R==NULL||L->val!=R->val){
            return false;
        }
        return isSame(L->left,R->right)&&isSame(L->right,R->left);
    }
    bool isSymmetric(TreeNode* root) {
        if(root==NULL){
            return true;
        }
        return isSame(root->left,root->right);
    }
};
