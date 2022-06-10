class Solution {
public:
    string replaceSpace(string s) {
        int cap = s.size();
        int count = 0;
        for(char ch : s)
        {
            if(ch == ' ') count++;
        }
        s.resize(cap+2*count);
        int j = s.size()-1;
        for(int i=cap-1;i>=0;i--)
        { 
            cout << s[i]  << " " << j <<endl;
            if(s[i]!=' ') s[j--] = s[i]; 
            else 
            {

                    s[j--] = '0';
                    s[j--] = '2';
                    s[j--] = '%';
            }
            
        }
        return s;
        
    }
};