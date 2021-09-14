class Solution {
public:
    int majorityElement(vector<int>& nums) {
        int candiate,count=0;
        for(int i=0;i<nums.size();i++){
            if(count==0){
                candiate=nums[i];
            }
            if(candiate==nums[i]){
                count++;
            }else{
                count--;
            }
        }
        return candiate;
    }
};
