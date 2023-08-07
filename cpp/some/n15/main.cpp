#include <vector>
#include <iostream>

using namespace std;

// TLE
class Solution final {
public:
    int minCostClimbingStairs(vector<int>& cost) {
        const auto n = cost.size();
        return std::min(minCost(cost, n - 1),
                        minCost(cost, n - 2));
    }

    int minCost(std::vector<int>& cost, size_t n) {
        if (n < 2) return cost[n];
        return cost[n] + std::min(minCost(cost, n - 1),
                                  minCost(cost, n - 2));
    }
};

class Solution1 final {
public:
    int minCostClimbingStairs(const vector<int>& cost) {
        const size_t n = cost.size();
        vector<int> dp; dp.resize(n, 0);

        for (size_t i = 0; i < n; ++i) {
            if (i < 2) {
                dp[i] = cost[i];
            }
            else {
                dp[i] = cost[i] +
                        std::min(dp[i - 1], dp[i - 2]);
            }
        }
        return std::min(dp[n - 1], dp[n - 2]);
    }
};

int main()
{
    cout << "Hello World!" << endl;
    return 0;
}
