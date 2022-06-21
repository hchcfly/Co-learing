//dfs
class Solution {
public:
    int book[105][105] = {0};

    int add(int x,int y)
    {
        int sum = 0;
        while(x)
        {
            sum+=(x%10);
            x /= 10;
        }
        while(y)
        {
            sum+=(y%10);
            y /= 10;
        }
        return sum;
    }

    void dfs(int m, int n,int i,int j,int k,int &res)
    {
        if(i<0 || i>=m || j<0 || j>=n || book[i][j] == 1 || add(i,j) > k)
        {
            return ;
        }

        book[i][j] = 1;
        res++;
        dfs(m,n,i,j+1,k,res);
        dfs(m,n,i+1,j,k,res);
    }


    
    int movingCount(int m, int n, int k) 
    {
        int res=0;
        dfs(m,n,0,0,k,res);
        return res;
    }
};
//dp
// class Solution {
// public:

//     int add(int x,int y)
//     {
//         int sum = 0;
//         while(x)
//         {
//             sum+=(x%10);
//             x /= 10;
//         }
//         while(y)
//         {
//             sum+=(y%10);
//             y /= 10;
//         }
//         return sum;
//     }

//     int movingCount(int m, int n, int k) 
//     {
//         vector<vector<bool>> dp(m+1,vector<bool>(n+1,false));
//         int res=0;
//         dp[0][1] = true;
//         for(int i=0;i<m;i++)
//         {
//             for(int j=0;j<n;j++)
//             {
//                 if((dp[i][j+1] || dp[i+1][j]) && add(i,j) <= k)
//                 {
//                     dp[i+1][j+1] = true;
//                     res++;
//                 }
//             }
//         }
//         return res;
//     }
// };

