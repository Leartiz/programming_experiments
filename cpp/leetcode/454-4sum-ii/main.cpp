// https://leetcode.com/problems/4sum-ii/

#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

class Solution
{
public:
    int fourSumCount(
        vector<int> &nums1,
        vector<int> &nums2,
        vector<int> &nums3,
        vector<int> &nums4)
    {
        // завести мапу, типо map<int, pair<int>>.
        // сумма + индексы. или вообще можно множество использовать.

        const int n = (int)nums1.size();

        unordered_map<int, int> sum_12; // сумма + количество таких сумм
        for (int i = 0; i < n; ++i)
        { // 200*200 максимальная сложность O(n*n)
            for (int j = 0; j < n; ++j)
            {
                sum_12[nums1[i] + nums2[j]] += 1; // все возможные суммы, первых двух массивов
            }
        }

        int result = 0;
        for (int i = 0; i < n; ++i)
        { // O(n*n)
            for (int j = 0; j < n; ++j)
            {

                // NOTE:
                //   nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
                //   nums1[i] + nums2[j] == -(nums3[k] + nums4[l])

                const auto sum_34 = nums3[i] + nums4[j];
                result += sum_12[-sum_34];
            }
        }

        return result;
    }
};

int main()
{
    Solution s;
    {
        std::vector<int> nums1{1, 2};
        std::vector<int> nums2{-2, -1};
        std::vector<int> nums3{-1, 2};
        std::vector<int> nums4{0, 2};
        std::cout << s.fourSumCount(nums1, nums2, nums3, nums4) << '\n'; // expect 2
    }
    {
        std::vector<int> nums1{-1, -1};
        std::vector<int> nums2{-1, 1};
        std::vector<int> nums3{-1, 1};
        std::vector<int> nums4{1, -1};
        std::cout << s.fourSumCount(nums1, nums2, nums3, nums4) << '\n'; // expect 6
    }
}
