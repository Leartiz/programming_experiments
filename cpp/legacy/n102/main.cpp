#include <iostream>

class A {
public:
    A() {
        A::print();
    }

    virtual void print() {
        std::cout << "A()\n";
    }
};

class B : public A {
public:
    B() {
        B::print();
    }

    void print() override {
        std::cout << "B()\n";
    }
};

int main()
{
    B();
    return 0;
}
