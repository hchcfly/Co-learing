/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:

    TreeNode* CreateTree(vector<int>& preorder,vector<int>& inorder)
    {
        //前序遍历大小为0
        if(preorder.size() == 0) return nullptr;

        //获取根节点信息
        int rootval = preorder[0];
        TreeNode* root = new TreeNode(rootval);
        if(preorder.size() == 1) return root;

        // find index
        int i=0;
        for(i=0;i<inorder.size();i++)
        {
            if(rootval == inorder[i])
            {
                break;
            }
        }

        //划分数组
        //中
        vector<int> inorder_left(inorder.begin(),inorder.begin()+i);
        vector<int> inorder_right(inorder.begin()+i+1,inorder.end());
        //先
        vector<int> preorder_left(preorder.begin()+1,preorder.begin()+1+inorder_left.size());
        vector<int> preorder_right(preorder.begin()+inorder_left.size()+1,preorder.end());
        root->left = CreateTree(preorder_left,inorder_left);
        root->right = CreateTree(preorder_right,inorder_right);
        return root;
    }

    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) 
    {
        return CreateTree(preorder,inorder);
    }
};