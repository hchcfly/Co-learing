class Solution 
{
public:
    int fib(int n) 
    {
        //0 1 1 2
        int a=0;
        int b=1;
        int num = n;
        int sum=0;
        while(num--)
        {
            sum=(a+b)%1000000007;
            a = b;
            b = sum;
        }
        return a;
    }
};