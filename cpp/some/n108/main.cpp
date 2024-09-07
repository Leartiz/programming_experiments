<<<<<<< HEAD
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
=======
#include <stdio.h>
#include <winsock2.h>

#define PORT 18080

int main() {

    WSADATA wsa;
    SOCKET server_socket, client_socket;

    struct sockaddr_in server, client;
    int addrlen = sizeof(server);
    char hello[] = "hi man!"; // Отправляем один байт на клиента

    // Initialize Winsock
    if (WSAStartup(MAKEWORD(2, 2), &wsa) != 0) {
        printf("WSAStartup failed\n");
        return 1;
    }

    // Create a socket
    if ((server_socket = socket(AF_INET, SOCK_STREAM, 0)) == INVALID_SOCKET) {
        printf("Socket creation failed\n");
        return 1;
    }

    server.sin_family      = AF_INET;
    server.sin_addr.s_addr = INADDR_ANY;
    server.sin_port        = htons(PORT);

    // Bind the socket
    if (bind(server_socket, (struct sockaddr*)&server, sizeof(server)) == SOCKET_ERROR) {
        printf("Bind failed\n");
        return 1;
    }

    // Listen for incoming connections
    if (listen(server_socket, 3) == SOCKET_ERROR) {
        printf("Listen failed\n");
        return 1;
    }

    printf("Server started, waiting for incoming connections...\n");

    while(true) {
        // Accept incoming connection
        if ((client_socket = accept(server_socket, (struct sockaddr*)&client, &addrlen)) == INVALID_SOCKET) {
            printf("Accept failed\n");
            return 1;
        }

        // Send data to client
        send(client_socket, hello, sizeof(hello), 0);
        printf("Byte successfully sent to client\n");

        closesocket(client_socket);
    }

    closesocket(server_socket);
    WSACleanup();

>>>>>>> d26a70116fb70142de5c81e6998d458a135dc7b7
    return 0;
}
