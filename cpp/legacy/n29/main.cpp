#include <iostream>
#include <algorithm>
#include <cassert>

using namespace std;

class Solution {
public:
    string toHex(int num) {
        if (!num) return { "0" };

        string result;
        const string hexs =
            "0123456789abcdef";

        auto positive_num =
            static_cast<uint32_t>(num);

        while (positive_num) {
            const int nibble =
                positive_num & 0xf;
            result.push_back(hexs[nibble]);
            positive_num = (positive_num >> 4);
        }

        reverse(result.begin(), result.end());
        return result;
    }
};

int main()
{
    Solution s;
    {
        assert(s.toHex(0) == "0");
        assert(s.toHex(1) == "1");
        assert(s.toHex(10) == "a");
        assert(s.toHex(-1) == "ffffffff");
        assert(s.toHex(26) == "1a");
        std::cout << "[OK]\n";
    }

    return 0;
}
