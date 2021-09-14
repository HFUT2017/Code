class Solution {
public: 
    int height(TreeNode* root){
        if(root==NULL){
            return 0;
        }
        int left=height(root->left);
        int right=height(root->right);
        return max(left,right)+1;
    }
    bool isBalanced(TreeNode* root) {
        if(root==NULL){
            return true;
        }
        int left=height(root->left);
        int right=height(root->right);
        if(max(left,right)-min(left,right)>1){
            return false;
        }
        return isBalanced(root->left)&&isBalanced(root->right);
    }
};
