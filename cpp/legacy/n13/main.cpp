#include <vector>
#include <iostream>
#include <iterator>
#include <cassert>

using namespace std;

// without extra space.
class Solution final{
public:
    vector<int> findDisappearedNumbers(vector<int>& nums) {
        for (size_t i = 0; i < nums.size(); ++i) {
            if (!nums[i]) continue;
            size_t curIndx = nums[i] - 1;

            while (nums[curIndx]) {
                const size_t tmp = curIndx;
                curIndx = nums[curIndx] - 1;
                nums[tmp] = 0;
            }
        }

        vector<int> result;
        for (size_t i = 0; i < nums.size(); ++i) {
            if (nums[i])
                result.push_back(i + 1);
        }

        return result;
    }
};

// O(n) and without extra space.
class Solution1 final{
public:
    vector<int> findDisappearedNumbers(vector<int>& nums) {
        for (size_t i = 0; i < nums.size(); ++i) {
            const size_t curIndx = std::abs(nums[i]) - 1;
            if (nums[curIndx] > 0) nums[curIndx] *= -1;
        }

        vector<int> result;
        for (size_t i = 0; i < nums.size(); ++i) {
            if (nums[i] > 0) result.push_back(i + 1);
        }

        return result;
    }
};

// -----------------------------------------------------------------------

void printlnVec(const std::vector<int>& vec) {
    std::copy(begin(vec), end(vec),
              std::ostream_iterator<int>(std::cout, " "));
    std::cout << std::endl;
}

int main()
{
    Solution s;
    {
        vector<int> nums = { 4, 3, 2, 7, 8, 2, 3, 1 };
        const auto got = s.findDisappearedNumbers(nums);
        printlnVec(got);
        assert(s.findDisappearedNumbers(nums) == std::vector({ 5, 6 }));
    }

    std::cout << "[OK]\n";
    return 0;
}
