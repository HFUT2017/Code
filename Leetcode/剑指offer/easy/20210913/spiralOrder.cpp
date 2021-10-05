class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
            vector<int>res;
            int n=matrix.size();
            if(n==0){
                return res;
            }
            int m=matrix[0].size();
            int left=0,right=m-1,top=0,bottom=n-1;
            while(left<=right&&top<=bottom){
                for(int i=left;i<=right;i++){
                    res.push_back(matrix[top][i]);
                }
                for(int i=top+1;i<=bottom;i++){
                    res.push_back(matrix[i][right]);
                }
                if(left<right&&top<bottom){
                    for(int i=right-1;i>=left;i--){
                        res.push_back(matrix[bottom][i]);
                    }
                    for(int i=bottom-1;i>top;i--){
                        res.push_back(matrix[i][left]);
                    }
                }
                left++;
                right--;
                top++;
                bottom--;
            }
            return res;
    }
};
