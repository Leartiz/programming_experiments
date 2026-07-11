// https://leetcode.com/problems/3sum-with-multiplicity/

#include <iostream>
#include <vector>
#include <array>
#include <algorithm>

#include "combinations.hpp"

using namespace std;

class Solution
{
public:
    int threeSumMulti(vector<int> &arr, const int target)
    {
        // 1.
        //
        // arr[i] + arr[j] + arr[k] == target
        //

        // 2.
        //
        // 0 <= arr[i] <= 100
        // сложить в массив частот
        //

        constexpr auto maxArrSize = 101;
        array<int, maxArrSize> freqs;
        freqs.fill(0);

        // O(n)
        for (int i = 0; i < arr.size(); ++i)
        {
            const auto indx = arr[i];
            freqs[indx] += 1;
        }

        sort(arr.begin(), arr.end());
        const auto last = unique(arr.begin(), arr.end());
        arr.erase(last, arr.end());

        long long result64 = 0;
        for (int i = 0; i < arr.size(); ++i) // [0, 100]
        {
            for (int j = i; j < arr.size(); ++j)
            {
                const auto a = arr[i];
                const auto b = arr[j];

                const auto c = target - (a + b);
                if (c < b || c > 100)
                {
                    continue;
                }

                if (a == b && b == c)
                {
                    result64 += combinations(freqs[a], 3);
                }
                else if (a == b)
                {
                    result64 += combinations(freqs[a], 2) * freqs[c];
                }
                else if (b == c)
                {
                    result64 += combinations(freqs[b], 2) * freqs[a];
                }
                else
                {
                    result64 += (freqs[a] * freqs[b] * freqs[c]);
                }
            }
        }

        return int(result64 % (1000'000'000 + 7));
    }
};

int main()
{
    Solution s;

    {
        std::vector<int> arr{1, 1, 2, 2, 3, 3, 4, 4, 5, 5};
        std::cout << s.threeSumMulti(arr, 8) << '\n'; // expect 20
    }
    {
        std::vector<int> arr{1, 1, 2, 2, 2, 2};
        std::cout << s.threeSumMulti(arr, 5) << '\n'; // expect 12
    }
}
