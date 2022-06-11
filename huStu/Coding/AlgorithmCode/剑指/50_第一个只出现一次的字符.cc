class Solution {
public:
    char firstUniqChar(string s) {
        unordered_map<char,bool> umap;
        vector<char> vc;
        for(auto ch : s)
        {
            if(umap.find(ch)==umap.end())
            {
                vc.push_back(ch);
                umap[ch] = false;
            }
            else
            {
                umap[ch] = true;
            }
        }
        for(auto ch : vc)
        {
            if(umap[ch] == false) return ch;
        }
        return ' ';
    }
};