class Solution {
public:
    TreeNode* mirrorTree(TreeNode* root) {
        if(root==NULL){
            return NULL;
        }
        swap(root->left,root->right);
        mirrorTree(root->left);
        mirrorTree(root->right);
        return root;
    }
};
