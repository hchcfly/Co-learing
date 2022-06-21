class Solution {
public: 
    //采用标记法
    bool dfs(vector<vector<char>>& board,string &word,int x,int y,int k)
    {
        //递归结束条件
        if(x < 0 || x > rows || y < 0 || y > cols || board[x][y] != word[k]) return false;
        //if(board[x][y] == '\0') return false; 和前一句逻辑重合
        if(k == word.size()-1)
        {
            return true;
        }
        //递归枚举条件
        board[x][y] = '\0';
        bool ans = dfs(board,word,x,y+1,k+1) || dfs(board,word,x+1,y,k+1) || dfs(board,word,x,y-1,k+1) || dfs(board,word,x-1,y,k+1);
        board[x][y] = word[k];
        return ans;
    }


    int rows,cols;
    
    bool exist(vector<vector<char>>& board, string word) 
    {
        rows = board.size()-1;
        cols = board[0].size()-1;
        for(int i=0;i<=rows;i++)
        {
            for(int j=0;j<=cols;j++)
            {
                if(dfs(board,word,i,j,0) == true) return true;
            }
        }
        return false;
    }
};