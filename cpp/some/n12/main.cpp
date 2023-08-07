#include <vector>
#include <iostream>
#include <cassert>

using namespace std;

// O(n)
class Solution final {
public:
    int maxProfit(const vector<int>& prices) {
        auto mxp = 0;
        auto buy = prices.front();
        for (size_t i = 1; i < prices.size(); ++i) {
            if (prices[i] > buy) {
                mxp = std::max(mxp, prices[i] - buy);
            }
            else {
                buy = prices[i];
            }
        }
        return mxp;
    }
};

// TLE
class Solution1 final {
public:
    int maxProfit(const vector<int>& prices) {
        int mxp{ 0 };
        for (size_t i = 0; i < prices.size(); ++i) {
            for (size_t j = i + 1; j < prices.size(); ++j) {
                const auto p = prices[j] - prices[i];
                if (mxp < p) mxp = p;
            }
        }
        return mxp;
    }
};

int main()
{
    Solution1 s1;
    {
        assert(s1.maxProfit({ 7, 1, 5, 3, 6, 4 }) == 5);
        assert(s1.maxProfit({ 7, 6, 4, 3, 1 }) == 0);
    }

    Solution1 s;
    {
        assert(s.maxProfit({ 7, 1, 5, 3, 6, 4 }) == 5);
        assert(s.maxProfit({ 7, 6, 4, 3, 1 }) == 0);
    }

    std::cout << "[OK]\n";
    return 0;
}
