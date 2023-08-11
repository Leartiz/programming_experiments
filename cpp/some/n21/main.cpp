#include <cmath>
#include <iostream>

using namespace std;

class Solution final
{
public:
    bool isPowerOfThree(int n) const {
        return n > 0 && 1162261467 % n == 0;
    }
};

void printAllPowers() {
    int power{ 0 };
    while (true) {
        const auto res{ static_cast<int>(std::pow(3, power)) };
        std::cout << power++ << " - " << res << std::endl;
        if (0 > res) break;
    }
}

int main()
{
    printAllPowers();
    return 0;
}
