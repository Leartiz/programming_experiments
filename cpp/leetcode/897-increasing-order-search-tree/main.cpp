// https://leetcode.com/problems/increasing-order-search-tree/

#include <iostream>

using namespace std;

struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution
{
public:
    TreeNode *increasingBST(TreeNode *root)
    {
        auto fake = new TreeNode;
        auto prev = fake; // копируем указатель
        toList(root, prev);
        return fake->right;
    }

    // NOTE: пример [5,1,7]
    /*
              5
             / \
            1   7

        inorder: 1, 5, 7

        start:  prev = fake,  chain: fake

        step1 visit 1:
            fake.right = 1,  prev = 1
            chain: fake -> 1

        step2 visit 5:
            1.right = 5,  prev = 5
            chain: fake -> 1 -> 5

        step3 visit 7:
            5.right = 7,  prev = 7
            chain: fake -> 1 -> 5 -> 7

        return fake->right  =>  1 -> 5 -> 7
    */
    void toList(TreeNode *node, TreeNode *&prev)
    {
        if (node == nullptr)
        {
            return;
        }

        toList(node->left, prev);
        node->left = nullptr;
        prev->right = node;
        prev = node;
        toList(node->right, prev);
    }
};

void printRightSpine(TreeNode *root)
{
    while (root != nullptr)
    {
        cout << root->val;
        if (root->right != nullptr)
        {
            cout << " -> ";
        }
        root = root->right;
    }
    cout << endl;
}

int main()
{
    Solution s;
    {
        // root = [5,1,7]
        // Output: [1,null,5,null,7]
        TreeNode *root = new TreeNode(5, new TreeNode(1), new TreeNode(7));
        printRightSpine(s.increasingBST(root));
    }
    {
        // root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
        // Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
        TreeNode *n1 = new TreeNode(1);
        TreeNode *n2 = new TreeNode(2, n1, nullptr);
        TreeNode *n4 = new TreeNode(4);
        TreeNode *n3 = new TreeNode(3, n2, n4);
        TreeNode *n7 = new TreeNode(7);
        TreeNode *n9 = new TreeNode(9);
        TreeNode *n8 = new TreeNode(8, n7, n9);
        TreeNode *n6 = new TreeNode(6, nullptr, n8);
        TreeNode *root = new TreeNode(5, n3, n6);
        printRightSpine(s.increasingBST(root));
    }
    return 0;
}
