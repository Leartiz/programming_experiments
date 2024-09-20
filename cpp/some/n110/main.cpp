#include <iostream>

struct A {
    void foo() {
        std::cout << "A::foo()"
                  << std::endl;
    }

    virtual void bar() {
        std::cout << "A::bar()"
                  << std::endl;
    }

    virtual ~A(){}
};

struct B : A {
    void foo() {
        std::cout << "B::foo()"
                  << std::endl;
    }

    void bar() override {
        std::cout << "B::bar()"
                  << std::endl;
    }
};

constexpr int inc (int a)
{
    return a + 1;
}

void func()
{
    constexpr int a = inc(3);

    // ошибка: a не является constexpr-выражением,
    // из-за чего возвращаемое значение не имеет спецификатор constexpr
    constexpr int b = inc(a);
}

struct b __attribute__ ((section ("bss")));

int main()
{
    func();

    // ***

    A a;
    B b;

    a.foo();
    b.foo();

    // ***

    a.bar();
    b.bar();

    A *pa = new B;
    pa->foo(); // !
    pa->bar();
    delete pa;

    // ***

    return 0;
}
