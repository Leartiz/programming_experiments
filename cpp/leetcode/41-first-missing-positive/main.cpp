// https://leetcode.com/problems/first-missing-positive/

#include <iostream>
#include <vector>

using namespace std;

class Solution
{
public:
    int firstMissingPositive(vector<int> &nums)
    {
        for (int i = 0; i < nums.size(); ++i)
        {
            while (true)
            {
                const auto val = nums[i]; // алиас
                if (val <= 0)
                {
                    break;
                }

                if (val > nums.size())
                { // 5 будет на позиции 5-1, а значит ОК
                    break;
                }

                if (nums[val - 1] == val)
                {
                    break;
                }

                swap(nums[i], nums[val - 1]);
            }
        }

        for (int i = 1; i <= nums.size(); ++i)
        {
            if (nums[i - 1] != i)
            {
                return i;
            }
        }

        return nums.size() + 1;
    }
};

int main()
{
    Solution s;
    {
        vector<int> nums = {3, 4, -1, 1};
        cout << s.firstMissingPositive(nums) << endl;
    }
    {
        vector<int> nums = {1, 2, 0};
        cout << s.firstMissingPositive(nums) << endl;
    }
    {
        vector<int> nums = {7, 8, 9, 11, 12};
        cout << s.firstMissingPositive(nums) << endl;
    }
    return 0;
}
