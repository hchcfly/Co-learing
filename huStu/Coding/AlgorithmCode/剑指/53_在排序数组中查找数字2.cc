class Solution {
public:
    int missingNumber(vector<int>& nums) 
    {
        //--------------------------10(N)
        // int len = nums.size();
        // for(int i=0;i<nums.size();i++)
        // {
        //     if(i!=nums[i]) return i;
        // }
        // return nums[len-1]+1;

        //--------------------------2二分0(logN)
        int left=0;
        int right=nums.size()-1;
        while(left < right)
        {
            int mid = (right-left)/2+left;
            if(nums[mid] == mid)
            {
                left = mid+1;
            }
            else
            {
                right = mid;
            }
        }
        cout << right << endl;
        return right==nums[nums.size()-1]?right+1:right;
    }
};