class Solution {
public:
    int integerBreak(int n) 
    {
        vector<int> dp(n+1);
        //动规
        //dp[i]拆分数字ｉ所能得到的最大乘积值
        //dp[i] = max(dp[i],max(j*(i-j),j*dp[i-j]))
        dp[2] = 1;
        for(int i=3;i<=n;i++)
        {
            //不拆分0 1
            for(int j=1;j<=i-1;j++)
            {
                dp[i] = max(dp[i],max(j*(i-j),j*dp[i-j]));
            }
        }
        return dp[n];
    }
};