// https://leetcode.com/problems/partition-array-into-three-parts-with-equal-sum/

#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    bool canThreePartsEqualSum(const vector<int>& arr) {
        // NOTE:
        /*
            len = [3, 50000]
            a[i] = [-10000, 10000]

            [0,2,1,-6,6,-7,9,1,2,0,1]

            [-7,-6,0,0,1,1,1,2,2,6,9]

            #1

            short a[20000]

            a[-7] = 1
            a[-6] = 1
                a[0] = 2
                a[1] = 3 исключить
                a[2] = 2
                a[6] = 1
                a[9] = 1

            #2

            we can find indexes

            [-7,-6,0,0,1,1,1,2,2,6,9]

            i = 1
            j = 3

                0  1   3             N
            [-7,-6,0,0,1,1,1,2,2,6,9]

            [) [) [)
        */

        // "базовые" суммы трёх частей
        auto part_1 = arr[0];
        auto src_part_2 = arr[1];
        auto src_part_3 = \
            sumFromTo(arr, 2, arr.size()); // [)

        if (part_1 == src_part_2 &&
            src_part_2 == src_part_3) {
            return true;
        }

        for (int i = 1; i < arr.size(); ++i) {

            part_1 += arr[i];
            src_part_2 = arr[i+1];
            src_part_3 = sumFromTo(arr, i+2, arr.size());

            auto part_2 = src_part_2;
            auto part_3 = src_part_3;

            for (int j = i + 1; j < arr.size() - 1; ++j) {

                part_2 += arr[j];
                part_3 -= arr[j];

                std::cout << i << " " << j << std::endl;
                std::cout << "part_1: " << part_1 << std::endl;
                std::cout << "part_2: " << part_2 << std::endl;
                std::cout << "part_3: " << part_3 << std::endl;
                if (part_1 == part_2 && part_2 == part_3) {
                    return true;
                }
            }
        }

        /*
        [0,2,1,-6,6,-7,9,1,2,0,1]
        1.
            0
            2
            1,-6,6,-7,9,1,2,0,1 = ?

        2.
            0 + 2 = 2
            1
            -6,6,-7,9,1,2,0,1 = ?

            1,-6
            6,-7,9,1,2,0,1
        */

        return false;
    }

    int sumFromTo(const std::vector<int>& arr, int beg, int end) {
        int sum = 0;
        for (int i = beg; i < end; ++i) {
            sum += arr[i];
        }
        return sum;
    }
};

int main() {
    Solution s;
    std::cout << std::boolalpha
              << s.canThreePartsEqualSum({0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1})
              << '\n'; // expect true
}
