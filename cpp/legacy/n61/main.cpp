#include <iostream>

using namespace std;

struct S { // The type has to be polymorphic
    virtual void f();
};

struct A {
    void f() {}
};

int main()
{
    {
        S* p = nullptr;
        try {
            std::cout << typeid(*p).name() << '\n';
        } catch(const std::bad_typeid& e) {
            std::cout << e.what() << '\n';
        }
    }

    // ***

    {
        A* p = nullptr;
        try {
            std::cout << typeid(*p).name() << '\n';
        } catch(const std::bad_typeid& e) {
            std::cout << e.what() << '\n';
        }
    }
}
