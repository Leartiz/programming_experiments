#include <cassert>
#include <iostream>
#include <algorithm>

using namespace std;

class Solution {
public:
    bool isSubsequence(const string& s, const string& t) {
        size_t si{ 0 };
        std::for_each(begin(t), end(t),
                      [&s, &si](const auto ch) -> void {
            if (ch == s[si]) si += 1;
        });
        return si == s.size();
    }
};

int main()
{
    return 0;
}
