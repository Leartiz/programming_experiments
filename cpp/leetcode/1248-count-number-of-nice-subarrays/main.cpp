// https://leetcode.com/problems/count-number-of-nice-subarrays/

#include <iostream>
#include <vector>

using namespace std;

class Solution
{
public:
    static inline bool isOdd(int val)
    {
        return val % 2 != 0;
    }

    int numberOfSubarrays(vector<int> &nums, int k)
    {
        // размер массива меньше чем требуемое количество нечетных
        if (nums.size() < k)
        {
            return 0;
        }

        int result = 0;

        // индексы
        int l = 0;
        int r = k - 1;

        int oddCount = 0; // текущее количество нечетных
        for (int i = l; i <= r; ++i)
        {
            if (isOdd(nums[i]))
            {
                ++oddCount;
            }
        }

        // NOTE:
        /*
            0   2
            1,1,2,1,1
        */
        while (l < nums.size())
        {
            // надо расширять окно если меньше
            // уменьшать если больше

            if (oddCount > k)
            {
                // так как удаляем именно прошлый
                if (isOdd(nums[l]))
                {
                    --oddCount;
                }
                ++l;
                continue;
            }

            if (oddCount == k)
            {
                ++result;
            }

            // а тут наоборот, идём вперед (пробуем увеличить окно сначало)
            if (r < nums.size() - 1)
            {
                ++r;
                if (isOdd(nums[r]))
                {
                    ++oddCount;
                }
                continue;
            }

            if (isOdd(nums[l]))
            {
                --oddCount;
            }
            ++l;
        }

        return result;
    }
};

int main()
{
    Solution s;
    {
        // Input: nums = [2,2,2,1,2,2,1,2,2,2], k = 2
        /*
            2,2,2,1,2,2,1,2,2,2
            2,2,2,1,2,2,1,2,2
            2,2,2,1,2,2,1,2
            2,2,2,1,2,2,1
              2,2,1,2,2,1,2,2,2 ---> 5
                2,1,2,2,1,2,2,2
                  1,2,2,1,2,2,2
              2,2,1,2,2,1,2,2
                2,1,2,2,1,2,2
              2,2,1,2,2,1,2 ---> 10
              2,2,1,2,2,1
                  1,2,2,1,2,2
                2,1,2,2,1,2
                  1,2,2,1,2
                2,1,2,2,1 ---> 15
                  1,2,2,1 ---> 16
        */
        // Output: 16
        vector<int> nums = {2, 2, 2, 1, 2, 2, 1, 2, 2, 2};
        cout << s.numberOfSubarrays(nums, 2) << endl;
    }
    {
        // Input: nums = [1,1,2,1,1], k = 3
        // Output: 2
        vector<int> nums = {1, 1, 2, 1, 1};
        cout << s.numberOfSubarrays(nums, 3) << endl;
    }
    {
        // Input: nums = [2,4,6], k = 1
        // Output: 0
        vector<int> nums = {2, 4, 6};
        cout << s.numberOfSubarrays(nums, 1) << endl;
    }
    return 0;
}
