class Solution {
public:
    string reverseWords(string s) 
    {
        string ans;
        for(int j=s.size()-1;j>=0;j--)
        {
            if(s[j] == ' ') continue;

            int i=j;
            while(i>=0 && s[i]!=' ') i--;
            ans.append(s.substr(i+1,j-i));
            ans.append(" ");
            j = i;
        }
        ans.pop_back();
        return ans;
    }
};