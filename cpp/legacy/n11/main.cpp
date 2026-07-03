#include <cassert>

#include <iostream>
#include <iterator>

#include <vector>

using namespace std;

class Solution final {
public:
    std::vector<int> countBits(const int n) {
        std::vector<int> result;
        result.reserve(static_cast<size_t>(n + 1));
        for (int i = 0; i <= n; ++i)
            result.push_back(zeroCountInBinaryStr(i));
        return result;
    }

private:
    int zeroCountInBinaryStr(int value) {
        int result{ 0 };
        while (value) {
            if (value % 2)
                ++result;
            value /= 2;
        }
        return result;
    }
};

void printlnVec(const std::vector<int>& vec) {
    std::copy(begin(vec), end(vec),
              std::ostream_iterator<int>(std::cout, " "));

    std::cout << std::endl;
}

int main()
{
    Solution s;
    {
        const auto got = s.countBits(2); printlnVec(got);
        assert(got == std::vector({ 0, 1, 1 }));
    }
    {
        const auto got = s.countBits(5); printlnVec(got);
        assert(got == std::vector({ 0, 1, 1, 2, 1, 2 }));
    }
    {
        const auto got = s.countBits(10); printlnVec(got);
        assert(got == std::vector({ 0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2 }));
    }
    //...

    return 0;
}
