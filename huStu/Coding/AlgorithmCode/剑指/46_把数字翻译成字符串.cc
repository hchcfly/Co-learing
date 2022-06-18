class Solution {
public:
    int translateNum(int num) 
    {
        string str = to_string(num);
        //dp[i]以i结尾的字符串的翻译方法
        int len = str.size();
        vector<int> dp(len+1);
        //初始化
        dp[0] = 1;
        dp[1] = 1;
        for (int i = 2; i <= len; i++)
        {
            //str[i-2]str[i-1]整体可以被翻译
            if(str[i-2]=='1' || (str[i-2]=='2' && str[i-1] <= '5'))
            {
                dp[i] = dp[i - 1] + dp[i - 2];
            }
            else
            {
                dp[i] = dp[i - 1];
            }
        }
        return dp[len + 1];
    }
};