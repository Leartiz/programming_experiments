#include <vector>
#include <iostream>

#include <functional>
#include <numeric>
#include <limits>

#include <cassert>

using namespace std;

// TLE.
class NotSolution final {
public:
    vector<int> productExceptSelf(const vector<int>& nums) {
        vector<int> result;
        result.reserve(nums.size());

        for (size_t i = 0; i < nums.size(); ++i) {
            int product{ 1 };
            for (size_t j = 0; j < nums.size(); ++j) {
                if (i == j) continue;
                product *= nums[j];
            }
            result.push_back(product);
        }

        return result;
    }
};

// -----------------------------------------------------------------------

// with space
class Solution final {
public:
    vector<int> productExceptSelf(const vector<int>& nums) {
        const size_t nn = nums.size();

        // ***

        vector<int> leRes(nn);
        leRes.front() = 1;

        for (size_t i = 1; i < nn; ++i)
            leRes[i] = (leRes[i - 1] * nums[i - 1]);

        // ***

        vector<int> riRes(nn);
        riRes.back() = 1;

        for (size_t i = nn - 2; i != std::numeric_limits<size_t>::max(); --i)
            riRes[i] = (riRes[i + 1] * nums[i + 1]);

        // ***

        vector<int> result;
        result.reserve(nn);

        for (size_t i = 0; i < nn; ++i)
            result.push_back(leRes[i] * riRes[i]);

        return result;
    }
};

// no space
class Solution1 final {
public:
    vector<int> productExceptSelf(const vector<int>& nums) {
        vector<int> result(nums.size(), 1);
        int pre{ 1 }, suf{ 1 };
        for (size_t  i = 0, n = nums.size(); i < n; i++) {
            result[i] *= pre;
            pre *= nums[i];

            // ***

            result[n - i - 1] *= suf;
            suf *= nums[n - i - 1];
        }
        return result;
    }
};

// -----------------------------------------------------------------------

int main()
{
    {
        NotSolution s;
        assert(s.productExceptSelf({ 1, 2, 3 }) == vector<int>({ 6, 3, 2 }));
        assert(s.productExceptSelf({ 1, 2, 3, 4 }) == vector<int>({ 24, 12, 8, 6 }));
        assert(s.productExceptSelf({ -1, 1, 0, -3, 3 }) == vector<int>({ 0, 0, 9, 0, 0 }));
    }

    {
        Solution s;
        assert(s.productExceptSelf({ 1, 2, 3 }) == vector<int>({ 6, 3, 2 }));
        assert(s.productExceptSelf({ 1, 2, 3, 4 }) == vector<int>({ 24, 12, 8, 6 }));
        assert(s.productExceptSelf({ -1, 1, 0, -3, 3 }) == vector<int>({ 0, 0, 9, 0, 0 }));
    }

    {
        Solution1 s;
        assert(s.productExceptSelf({ 1, 2, 3 }) == vector<int>({ 6, 3, 2 }));
        assert(s.productExceptSelf({ 1, 2, 3, 4 }) == vector<int>({ 24, 12, 8, 6 }));
        assert(s.productExceptSelf({ -1, 1, 0, -3, 3 }) == vector<int>({ 0, 0, 9, 0, 0 }));
    }

    std::cout << "[OK]\n";
    return 0;
}
