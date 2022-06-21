class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) 
    {
        int len = nums.size();
        for(int i=0;i<len;i++)
        {
            int ans = target-nums[i];
            int left = i+1;
            int right = len-1;
            //二分查找
            while(left <= right)
            {
                int mid = left+(right-left)/2;
                if(nums[mid] == ans)
                {
                    return vector<int>{ans,nums[i]};
                }
                else if(nums[mid] < ans)
                {
                    left = mid+1;
                }
                else
                {
                    right = mid-1;
                }
            }
        }
        return nums;
    }
};