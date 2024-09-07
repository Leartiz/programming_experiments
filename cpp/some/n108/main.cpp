#include <stack>
#include <vector>
#include <iostream>
#include <cassert>
#include <limits>

using namespace std;

struct TreeNode {
    int val;

    TreeNode *left;
    TreeNode *right;

    TreeNode()
        : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x)
        : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right)
        : val(x), left(left), right(right) {}
};

// -----------------------------------------------------------------------

class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> result;

        std::vector<TreeNode*> levelNodes;
        if (root) levelNodes.push_back(root);

        while (!levelNodes.empty()) {
            std::vector<int> levelValues;
            std::vector<TreeNode*> nextLevelNodes;

            while(!levelNodes.empty()) {
                auto node = levelNodes.front();
                levelNodes.erase(levelNodes.begin());
                levelValues.push_back(node->val);

                if (node->left) nextLevelNodes.push_back(node->left);
                if (node->right) nextLevelNodes.push_back(node->right);
            }

            result.push_back(std::move(levelValues));
            levelNodes = std::move(nextLevelNodes);
        }

        return result;
    }

    bool isValidBST(TreeNode* root) {
        std::stack<TreeNode*> levelNodes;
        if (root) levelNodes.push(root);

        while (!levelNodes.empty()) {
            auto node = levelNodes.top();
            levelNodes.pop();

            if (node->left) {
                if (node->left->val > min)
                    return false;
                levelNodes.push(node->left);
            }
            if (node->right) {
                if (node->right->val <= max)
                    return false;
                levelNodes.push(node->right);
            }
        }
        return true;
    }
};

// -----------------------------------------------------------------------

int main()
{
    TreeNode* root = new TreeNode(
        1, nullptr, new TreeNode(1));

    Solution s;
    assert(s.isValidBST(root));

    std::cout << "[OK]" << std::endl;
    return 0;
}
