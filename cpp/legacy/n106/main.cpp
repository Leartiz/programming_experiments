#include <iostream>
#include <cassert>
#include <vector>
#include <stack>
#include <set>

using namespace std;

class Solution {
public:
    set<int> goodSet{ 2, 3, 5 };

    bool isUgly(const int n) {
        if (n <= 0)
            return false;
        if (n == 1)
            return true;

        bool result = false;
        stack<pair<int, int>> cases;
        pushToCases(n, cases);

        // ***

        while (!cases.empty()) {
            const auto cs = cases.top();
            cases.pop();

            result = result ||
                     ugly(cs.first, cs.second, cases);
        }
        return result;
    }

    bool ugly(int n, int base, stack<pair<int, int>>& cases) {
        if (n % base != 0)
            return false;
        if (goodSet.count(n))
            return true;

        // ***

        n = n / base;
        pushToCases(n, cases);
        return false;
    }

    void pushToCases(int n, stack<pair<int, int>>& cases) {
        for (auto it = goodSet.cbegin();
             it != goodSet.cend(); ++it)
            cases.push({n, *it});
    }
};

// -----------------------------------------------------------------------

int main()
{
    Solution s;
    vector<pair<int, bool>> tests{
        { 0,  false },
        { 6,   true },
        { 1,   true },
        { 14, false },
        //...
    };
    for (size_t i = 0; i < tests.size(); ++i) {
        cout << s.isUgly(tests[i].first) << endl;

        assert(s.isUgly(tests[i].first)
               == tests[i].second);
    }
    return 0;
}
