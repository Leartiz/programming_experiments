#include <string>
#include <vector>
#include <bitset>
#include <iostream>

using namespace std;

class Solution {
public:
    vector<string> readBinaryWatch(int num) {
        vector<string> rs;
        for (int h = 0; h < 12; h++)
            /*
                0 = 0000
                1 = 0001
                2 = 0010
                3 = 0011
                4 = 0100
                _
                5 = 0101
                6 = 0110
                7 = 0111
                8 = 1000
                _
                9  = 1001
                10 = 1010
                11 = 1011
                12 = 1100
            */
            for (int m = 0; m < 60; m++)
                if (bitset<10>(m << 4 | h).count() == size_t(num))
                    rs.emplace_back(to_string(h) + (m < 10 ? ":0" : ":") + to_string(m));
        return rs;
    }
};

void printVec(const std::vector<string>& vec) {
    for (size_t i = 0; i < vec.size(); ++i)
        std::cout << vec[i] << " ";
    std::cout << std::endl;
}

int main()
{
    Solution s;
    printVec(s.readBinaryWatch(1));
    // printVec(s.readBinaryWatch(2));
    // printVec(s.readBinaryWatch(3));
    // printVec(s.readBinaryWatch(4));
    return 0;
}
