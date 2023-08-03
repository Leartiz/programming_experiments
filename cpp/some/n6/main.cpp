#include <stack>
#include <iostream>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left{ nullptr };
    TreeNode *right{ nullptr };

    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    explicit TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution final
{
public:
    bool isSubtree(TreeNode* root, TreeNode* subRoot) {
        std::stack<TreeNode*> nodes;
        nodes.push(root);

        while (!nodes.empty()) {
            const auto top = nodes.top();
            nodes.pop();

            if (isSameStructure(top, subRoot))
                return true;

            // ***

            if (top->left) nodes.push(top->left);
            if (top->right) nodes.push(top->right);
        }

        return false;
    }

    bool isSameStructure(TreeNode* lhs, TreeNode* rhs) {
        if (lhs == rhs) return true;
        if (lhs == nullptr || rhs == nullptr)
            return false;

        if (lhs->val != rhs->val)
            return false;

        return isSameStructure(lhs->left, rhs->left) &&
                isSameStructure(lhs->right, rhs->right);
    }
};

int main()
{
    return 0;
}
