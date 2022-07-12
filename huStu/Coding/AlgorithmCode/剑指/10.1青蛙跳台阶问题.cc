class Solution {
public:
    int numWays(int n) 
    {
        //1 1 2 
        //n {n-1 n-2}
        //f(n) = f(n-1)+f(n-2)
        int a = 1;
        int b = 1;
        int sum=0;
        int num = n;
        while(num--)
        {
            sum = (a+b)%1000000007;
            a = b;
            b = sum;
        }
        return a;
    }
};