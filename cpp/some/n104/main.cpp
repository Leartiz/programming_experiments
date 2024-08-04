#include <iostream>
#include <vector>
#include <string>
#include <sstream>
#include <map>

using namespace std;

class Solution {
public:

    // Input: mapping = [8,9,4,0,2,1,3,5,7,6],
    //     nums = [991,338,38]
    // Output: [338,38,991]

    vector<int> sortJumbled(vector<int>& mapping,
                            vector<int>& nums) {
        std::map<int, std::vector<int>> mm;
        for (size_t i = 0; i < nums.size(); ++i) {
            auto strNum = std::to_string(nums[i]);
            for (size_t j = 0; j < strNum.size(); ++j)
                strNum[j] = mapping[strNum[j] - '0'] + '0';


            // ***

            int convertedVal = 0;
            std::istringstream ss(strNum); ss >> convertedVal;
            mm[convertedVal].push_back(nums[i]);
        }

        vector<int> results;
        results.reserve(nums.size());
        for (auto it = mm.begin(); it != mm.end(); ++it)
            std::copy(it->second.begin(), it->second.end(),
                      std::back_inserter(results));

        return results;
    }
};

int main()
{
    vector<int> mapping = { 8,9,4,0,2,1,3,5,7,6 };
    vector<int> nums = { 991,338,38 };

    Solution s;
    const auto result = s.sortJumbled(mapping, nums);
    for (size_t i = 0; i < result.size(); ++i)
        std::cout << result[i] << " ";

    return 0;
}
