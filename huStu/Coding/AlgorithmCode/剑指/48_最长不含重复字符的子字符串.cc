class Solution {
public:
    int lengthOfLongestSubstring(string s) 
    {
        bool book[256] = false;
        int left = 0;
        int right = 0;
        int len = s.size();
        int ans = 0;
        //滑动窗口
        while(right < len)
        {
            while(book[s[right]] == true)
            {
                book[s[left]] = false;
                left++;
            }
            book[s[right]] = true;
            ans = max(ans, right - left + 1);
            right++;
        }
        return ans;
    }
};