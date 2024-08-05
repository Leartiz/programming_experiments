#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

void printVec(const std::vector<int>& vec) {
    for (auto& val : vec)
        std::cout << val << " ";
    std::cout << "\n";
}

class Solution {
public:
    int numTeams(const vector<int>& rating) {

        auto copyRating = std::vector<int>(rating.size(), 0);
        std::reverse_copy(rating.begin(), rating.end(), copyRating.begin());

        auto f = [](const vector<int>& rating) -> int {
            auto ratingPath = std::vector<int>(2, 0);
            int resultCount = 0;

            for (size_t i = 0; i < rating.size(); ++i) {
                ratingPath[0] = rating[i]; // base case!

                for (size_t j = i + 1; j < rating.size(); ++j) {
                    if (ratingPath[0] > rating[j]) {
                        ratingPath[1] = rating[j];
                        for (size_t k = j + 1; k < rating.size(); ++k) {
                            if (ratingPath[1] > rating[k])
                                ++resultCount;
                        }
                    }
                }
            }
            return resultCount;
        };
        return f(rating)+f(copyRating);
    }
};

int main()
{
    Solution s;
    std::cout << s.numTeams({2,5,3,4,1}) << "\n";
    std::cout << s.numTeams({2,1,3}) << "\n";
    std::cout << s.numTeams({1,2,3,4}) << "\n";

    return 0;
}
