// https://leetcode.com/problems/number-of-valid-words-in-a-sentence/

#include <iostream>
#include <string>
#include <string_view>

using namespace std;

class Solution {
public:
    bool isValid(string_view sub) {
        auto curTok = sub.begin();
        if (curTok == sub.end()) {
            return false;
        }

        // check first
        if (*curTok == '!' || *curTok == ',' || *curTok == '.') {
            advance(curTok, 1);
            return curTok == sub.end();
        }
        if (*curTok == '-' || *curTok < 'a' || *curTok > 'z') {
            return false;
        }

        // check nexts
        while (curTok != sub.end()) {
            if (*curTok >= 'a' && *curTok <= 'z') {
                advance(curTok, 1);
                continue;
            }

            if (*curTok == '-') {
                advance(curTok, 1);
                if (curTok == sub.end()) {
                    return false;
                }
                if (*curTok == '!' || *curTok == '.' || *curTok == ',' || *curTok == '-') {
                    return false;
                }
                continue;
            }

            if (*curTok == '!' || *curTok == '.' || *curTok == ',') {
                advance(curTok, 1);
                if (curTok != sub.end()) {
                    return false;
                }
                continue;
            }

            return false;
        }

        return true;
    }

    int countValidWords(string sentence) {
        if (sentence.empty()) {
            return 0;
        }

        int c = 0;
        auto it = sentence.begin();
        decltype(it) begSubToken;

        while (it != sentence.end()) {
            if (*it == ' ') {
                advance(it, 1);
                continue;
            }

            begSubToken = it;

            while (it != sentence.end()) {
                if (*it == ' ') {
                    break;
                }
                advance(it, 1);
            }

            if (isValid(string_view(&*begSubToken, static_cast<size_t>(it - begSubToken)))) {
                ++c;
            }

            if (it != sentence.end()) {
                advance(it, 1);
            }
        }
        return c;
    }
};

int main() {
    Solution s;
    std::cout << s.countValidWords("cat dog cat") << '\n'; // expect 3
}