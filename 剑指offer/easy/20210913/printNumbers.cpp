class Solution {
public:
    vector<int> printNumbers(int n) {
        int res=0;
        vector<int>ans;
        for(int i=1;i<=n;i++){
            res=res*10+9;
        }
        for(int i=1;i<=res;i++){
            ans.push_back(i);
        }
        return ans;
    }
};
