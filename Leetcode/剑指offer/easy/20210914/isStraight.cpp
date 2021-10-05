class Solution {
public:
    bool isStraight(vector<int>& nums) {
        int Min=14,Max=-1;
        map<int,int>hash;
        for(int i=0;i<nums.size();i++){
            if(nums[i]==0){
                continue;
            }
            if(hash[nums[i]]!=0)
            return false;
            hash[nums[i]]=1;
            Max=max(Max,nums[i]);
            Min=min(Min,nums[i]);
        }
        return Max-Min<5;
    }
};
