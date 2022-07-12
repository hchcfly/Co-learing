class Solution {
public:
    int cuttingRope(int n) 
    {
        //拆分尽可能多的3,如果剩余一个4则不用拆分
        //2 1 1  1
        //3 1 2  2
        if(n<=3) return n-1;  
        long res = 1;
        while(n>4)
        {
            res = (res*3)%1000000007;
            n-=3;
        }

        return (res*n)%1000000007;
    }
};