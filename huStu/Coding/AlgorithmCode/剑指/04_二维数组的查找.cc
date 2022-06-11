class Solution {
public:
    bool findNumberIn2DArray(vector<vector<int>>& matrix, int target) 
    {
        if(matrix.size() == 0)
        {
            return false;
        }
        int i = 0;
        int j = matrix[0].size()-1;

        while(i>=0 && i<=matrix.size()-1 && j>=0 && j<=matrix[0].size()-1)
        {
            if(matrix[i][j] == target)
            {
                return true;
            }
            else if(matrix[i][j] > target)
            {
                j--;
            }
            else
            {
                i++;
            }
        }
        return false;


    }
};