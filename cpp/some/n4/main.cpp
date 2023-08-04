#include <cassert>
#include <iostream>
#include <vector>

class NotSolution {
public:
    int missingNumber(const std::vector<int>& nums) {
        int result{ 0 };

        // distinct numbers in the range [0, n]!
        for (size_t i = 0; i < nums.size(); ++i) {
            result ^= (nums[i] ^ i);
        }
        result ^= nums.size();
        return result;
    }
};

int main()
{
    {
        std::cout << (5 ^ 5) << std::endl;
        std::cout << (5 ^ 0) << std::endl;
        std::cout << (0 ^ 5) << std::endl;
        std::cout << (0b0011 ^ 0b0101) << " = " << 0b0110 << std::endl;
    }

    // tests
    NotSolution s;
    {
        assert(s.missingNumber({1, 2, 3}) == 0);
        // res         nums[i]      i
        //  |             |         |
        //  |             |         |
        //  V             V         V
        // 0b00   ^   ((0b0001 ^ 0b0000) = 0b01) = 0b01 = 1
        // 0b01   ^   ((0b0010 ^ 0b0001) = 0b11) = 0b10 = 2
        // 0b10   ^   ((0b0011 ^ 0b0010) = 0b01) = 0b11 = 3

        // 3 ^ 3 = 0
    }
    std::cout << ".";

    std::cout << "\n[OK]\n";
    return 0;
}
