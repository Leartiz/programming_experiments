#include <iostream>
#include <cstddef>

class Base1 {
public:
    int a;
    int a1;
};

class Base2 {
public:
    int b;
};

class Derived : public Base1, public Base2 {
public:
    char c;
};

int main() {
    std::cout << "Offset of Base1 in Derived: " << offsetof(Derived, a) << std::endl;
    std::cout << "Offset of Base1 in Derived: " << offsetof(Derived, a1) << std::endl;

    std::cout << "Offset of Base1 in Derived: " << offsetof(Derived, b) << std::endl;
    std::cout << "Offset of Base1 in Derived: " << offsetof(Derived, c) << std::endl;

    //std::cout << "Offset of Base2 in Derived: " << offsetof(Derived, Base2::b) << std::endl;
    //std::cout << "Offset of c in Derived: " << offsetof(Derived, c) << std::endl;

    return 0;
}
