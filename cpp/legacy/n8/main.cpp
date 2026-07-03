#include <map>
#include <string>
#include <vector>

#include <cassert>
#include <iostream>

using namespace std;

class Solution {
public:
    bool wordBreak(string s, const vector<string>& wordDict) {
        map<char, vector<pair<string, size_t>>> map; // word, start
        for (size_t i = 0; i < s.size(); ++i) {
            for (const auto& word : wordDict) {
                if (s[i] == word.front()) {
                    std::string sub; sub.reserve(word.size());
                    std::copy(s.begin() + i, s.begin() + i + word.size(),
                              back_inserter(sub));

                    if (word == sub)
                        map[s[i]].push_back({ sub, i });
                }
            }
        }


    }
};

int main()
{
    Solution s;
    {
        assert(s.wordBreak("leetcode", { "leet", "code" }));
        assert(!s.wordBreak("leetcode", { "leetco", "code" }));
    }

    return 0;
}
