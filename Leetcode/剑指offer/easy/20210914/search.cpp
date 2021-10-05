class Solution {
public:
    int search(vector<int>& nums, int target) {
        return helper(nums,target)-helper(nums,target-1);
    }
    int helper(vector<int>& nums,int target){
        int i=0,j=nums.size()-1;
        while(i<=j){
            int mid=(i+j)/2;
            if(nums[mid]<=target) i=mid+1;
            else
            j=mid-1;
        }
        return i;
    }
};
