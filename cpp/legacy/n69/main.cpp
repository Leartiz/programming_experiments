#include <iostream>

using namespace std;

void bar_ccpy(const int value) {
    cout << sizeof(value) << endl;
    cout << "ccpy value: " << value << endl;
}

void bar_cref(const int& value) {
    cout << sizeof(value) << endl;
    cout << "cref value: " << value << endl;
}

void cast_cref(const int& value) {

}

int main()
{
    {
        cout << "sizeof(const int*): " << sizeof(const int*) << endl;
        cout << "sizeof(const int&): " << sizeof(const int&) << endl;
        cout << "sizeof(const int): " << sizeof(const int) << endl;

        cout << "sizeof(int*): " << sizeof(int*) << endl;
        cout << "sizeof(int&): " << sizeof(int&) << endl;
        cout << "sizeof(int): " << sizeof(int) << endl;
    }

    // ***

    {
        int value = 101;
        bar_ccpy(value);
        bar_cref(value);
        cout << "value: " << value << endl;
    }

    return 0;
}
