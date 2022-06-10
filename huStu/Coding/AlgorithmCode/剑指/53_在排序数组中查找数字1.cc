class Solution {
public:
    int search(vector<int>& nums, int target) {
        //数组为空时的处理
        if(nums.size() == 0)
        {
            return 0;
        }
        //定义二分边界
        int left = 0;
        int right = nums.size()-1;
        //找到最左边的target值
        while(left < right)
        {
            int mid = (right-left)/2+left;
            if(nums[mid] >= target)
            {
                right = mid;
            }
            else
            {
                left = mid+1;
            }
        }
        int ans=0;
        //找到target则继续往后寻找相同值的个数
        if(nums[right] == target)
        {
            ans=1;
            for(int i=right;i<nums.size();i++)
            {
                if(i+1<nums.size() && nums[i]==nums[i+1])
                {
                    ans++;
                }
                else
                {
                    break;
                }
            }
        }
        return ans;
    }
};