class Solution {
public:
    int cuttingRope(int n) 
    {
        //对应于整数拆分
        //dp[i]表示:拆分数字i所能得到的乘积最大值
        //dp[i] = max(dp[i],max(j*(i-j),j*dp(i-j)));
        
        vector<int> dp(n+1);
        //初始化
        dp[2] = 1;
        for(int i=3;i<=n;i++)
        {
            for(int j=1;j<=i-2;j++)
            {
                dp[i] = max(dp[i],max(j*(i-j),j*dp[i-j]));
            }
        }
        return dp[n];

    }
};