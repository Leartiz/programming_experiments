#include <vector>
#include <string>
#include <iostream>

using namespace std;

class Solution 
{
private:
    bool isSorted(const vector<int>& powers, 
                  const string& lhs, const string& rhs) 
    {
        for (size_t i = 0; i < lhs.size(); ++i) {
            if (i >= rhs.size()) {
                return false;
            }

            const auto lhsCurChPower = powers[lhs[i] - 'a'];
            const auto rhsCurChPower = powers[rhs[i] - 'a'];

            if (lhsCurChPower < rhsCurChPower) {
                return true;
            } 
            else if (lhsCurChPower > rhsCurChPower) {
                return false;
            }
        }
        return true;
    }

public:
    bool isAlienSorted(const vector<string>& words, string order)
    {
        vector<int> powers(26, 0); // abcde...

        for (size_t i = 0; i < order.size(); ++i) {
            powers[order[i] - 'a'] = i;
        }

        for (size_t i = 0; i < words.size() - 1; ++i) {
            if (!isSorted(powers, words[i], words[i + 1])) {
                return false;
            }
        }
        return true;
    }
};

int main() 
{
    Solution s;
    {
        std::cout << std::boolalpha << s.isAlienSorted(
            { "hello", "leetcode" }, 
            "hlabcdefgijkmnopqrstuvwxyz"
        ) << std::endl;

        std::cout << std::boolalpha << s.isAlienSorted(
            { "word", "world", "row" }, 
            "worldabcefghijkmnpqstuvxyz"
        ) << std::endl;

        std::cout << std::boolalpha << s.isAlienSorted(
            { "apple", "app" }, 
            "abcdefghijklmnopqrstuvwxyz"
        ) << std::endl;
    }
    return 0;
}