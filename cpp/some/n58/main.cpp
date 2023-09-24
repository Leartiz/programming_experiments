#include <iostream>

using namespace std;

struct UnwShow final
{
    int number{ 0 };

    UnwShow() = default;
    explicit UnwShow(int number)
        : number{ number }
    {
        cout << "ctor: " << number << "\n";
        cout << "addr: " << this << endl;
    }
    ~UnwShow()
    {
        cout << "dtor: " << number << "\n";
        cout << "addr: " << this << endl;
    }
};

void foo(int n)
{
    UnwShow s{ n };

    //if (n == 0) abort();

    if (n == 0) {
        cout << "\tthrow\n";
        throw 1;
    }

    foo(n - 1);
}

int main()
{
    try {
        foo(5);
    }
    catch(int) {

        // stack unwinding...

        cout << "\tcatch\n";
    }
}
