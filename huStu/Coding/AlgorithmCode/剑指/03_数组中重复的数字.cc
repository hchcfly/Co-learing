class Solution {
public:
    // 原地构造方法
    int findRepeatNumber(vector<int>& nums) {
        int i=0;
        for(i=0;i<nums.size();)
        {
            if(nums[i]!=i)
            {
                if(nums[nums[i]]!=nums[i])  
                {
                    swap(nums[nums[i]],nums[i]);
                }
                else 
                {
                    return nums[i];
                }
            }
            else
            { 
                //只有已经定位了的元素才可以i++
                i++;
            }
        }
        return -1;
    }
};