#include <iostream>
#include <stdexcept>

using namespace std;

struct S
{
    int value{ 0 };

    ~S()
    {
        throw runtime_error("to terminate");
    }
};

int main()
{
    {
        try {
            S s{ 25 };
            cout << s.value << endl;
        }
        catch (...) {
            cerr << "err" << endl;
        }
    }

    return 0;
}
