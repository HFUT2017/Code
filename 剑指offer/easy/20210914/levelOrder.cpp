class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
   vector<vector<int>> res;
   if(root== nullptr){
       return res;
   }
   queue<TreeNode*> q;
   q.push(root);
   while(!q.empty()){
       int n=q.size();
       vector<int> ans;
       for(int i=0;i<n;i++){
           TreeNode *node=q.front();
           q.pop();
           ans.push_back(node->val);
           if(node->left!= nullptr){
               q.push(node->left);
           }
           if(node->right!= nullptr){
               q.push(node->right);
           }
       }
       res.push_back(ans);
   }
   return res;
}
};
