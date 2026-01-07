#include <vector>
#include <stack>
#include <string>
#include <string_view>

#include <iostream>
#include <iomanip>

using namespace std;

class Solution final
{
public:
    vector<string> restoreIpAddresses(const string& s)
    {
        basicString = s;
        ipAddrs = {};

        if (s.size() < 4)
            return std::move(ipAddrs);

        createNextIpAddrs("", { s[0] }, 0, { s.data() + 1 });
        createNextIpAddrs("", { s[0], s[1] }, 0, { s.data() + 2 });
        createNextIpAddrs("", { s[0], s[1], s[3] }, 0, { s.data() + 3 });

        return std::move(ipAddrs);
    }

private:
    void createNextIpAddrs(string curState, string part, int partNum, string_view rest)
    {
        //
            cout << curState << " " << part << " " << partNum << " " << rest << endl;
        //

        if (!isValidPartIp(part)) {
            return;
        }

        curState.append(std::move(part));
        if (partNum == 4 && rest.empty()) {
            ipAddrs.push_back(std::move(curState));
            return;
        }

        if (true) {

        }

        createNextIpAddrs(curState, string{ s[0] }, 0, string_view(s.data() + 1));
        createNextIpAddrs(curState, string{ s[0], s[1] }, 0, string_view(s.data() + 2));
        createNextIpAddrs(curState, string{ s[0], s[1], s[3] }, 0, string_view(s.data() + 3));
    }

    bool isValidPartIp(const string_view part)
    {

        return true;
    }

    string_view basicString;
    vector<string> ipAddrs;
};



int main()
{
    {
        const string input{ "25525511135" };
        {
            const string_view sv{ input.data(), 5 };
            cout << "sv.size: " << sv.size() << endl;
            cout << "sv.max_size: " << sv.max_size() << endl;
            cout << "sv.empty: " << sv.empty() << endl;
            cout << "sv.data: " << sv.data() << endl;
            cout << "sv: " << sv << endl;
        }
        {
            const string_view sv{ input };
            cout << "sv.size: " << sv.size() << endl;
            cout << "sv.max_size: " << sv.max_size() << endl;
            cout << "sv.empty: " << sv.empty() << endl;
            cout << "sv.data: " << sv.data() << endl;
            cout << "sv: " << sv << endl;
        }
        {
            const string_view sv{ input.data() + 11};
            cout << "sv.size: " << sv.size() << endl;
            cout << "sv.max_size: " << sv.max_size() << endl;
            cout << "sv.empty: " << sv.empty() << endl;
            cout << "sv.data: " << sv.data() << endl;
            cout << "sv: " << sv << endl;
        }
    }
    Solution s;
    {
        const string input{ "25525511135" };
        s.restoreIpAddresses(input);
    }
    {
        const string input{ "0000" };
        s.restoreIpAddresses(input);
    }
}

// Input: s = "25525511135"
// Output: ["255.255.11.135","255.255.111.35"]

// Input: s = "0000"
// Output: ["0.0.0.0"]

// Input: s = "101023"
// Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
