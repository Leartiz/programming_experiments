#include <vector>
#include <iostream>
#include <algorithm>

using namespace std;

vector<vector<int>> adj;


// function to add edge to the graph
void addEdge(int x,int y)
{
    adj[x][y] = 1;
    adj[y][x] = 1;
}

// Function to perform BFS on the graph
void bfs(int start, int end)
{
    const size_t n = adj.size();

    vector<int> d(n), p(n);
    p[start] = -1;
    vector<bool> used(n, false);

    vector<int> q;
    q.push_back(start);

    // Set source as visited
    used[start] = true;

    int vis;
    while (!q.empty()) {
        vis = q.front();
        q.erase(q.begin());

        // ***

        for (size_t i = 0; i < adj[vis].size(); ++i) {
            auto to = adj[vis][i];
            if (to == 1 && !used[i]) {
                q.push_back(i);
                used[i] = true;

                d[i] = d[vis] + 1;
                p[i] = vis;
            }
        }
    }

    std::cout << "***";

    vector<int> path;
    for (int v = end; v != -1; v = p[v])
        path.push_back(v);

    reverse(path.begin(), path.end());
    cout << "Path: ";
    for (size_t i=0; i<path.size(); ++i)
        cout << path[i] + 1 << " ";
}

// Driver code
int main()
{
    // number of vertices
    size_t v = 16;


    // adjacency matrix
    adj= vector<vector<int>>(v,vector<int>(v, 0));
    for (size_t i = 0; i < v; ++i) {
        adj[i][i] = 0;
    }

    addEdge(2, 3);
    addEdge(3, 7);
    addEdge(5, 9);
    addEdge(7, 11);

    addEdge(8, 9);
    addEdge(9, 10);
    addEdge(10, 11);

    addEdge(8, 12);
    addEdge(9, 10);
    addEdge(10, 14);
    addEdge(11, 15);

    addEdge(12, 13);
    addEdge(13, 14);
    addEdge(14, 15);

    for (size_t i = 0; i < v; ++i) {
        for (size_t j = 0; j < v; ++j) {
            std::cout << adj[i][j] << " ";
        } std::cout << std::endl;
    }

    // perform bfs on the graph
    bfs(2, 12);
}
