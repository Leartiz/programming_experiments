#include <algorithm>
#include <iostream>
#include <vector>
#include <string>

using namespace std;

class Solution {
public:
    int minDeletion(const string& s, int k) {
        vector<int> v(26, 0);

        for (size_t i = 0; i < s.size(); ++i) {
            ++v[size_t(s[i]) - 'a'];
        }

        v.erase(remove(v.begin(), v.end(), 0), v.end());
        sort(v.begin(), v.end(), greater<int>()); 

        return v.size();

        int result = 0;
        while (v.size() > k) {
            result += v.back();
            v.pop_back();
        }

        return result;
    }
};

int main() {
    Solution s;
    std::cout << s.minDeletion("yyyzz", 1) << std::endl;
}