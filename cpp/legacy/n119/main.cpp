#include <iostream>
#include <variant>
#include <any>

class Foo {
public:
    int a;
    int b;
};

int main()
{
    {
        std::variant<int, double, Foo> v;
        std::cout << "valueless_by_exception: " << v.valueless_by_exception() << '\n';
        std::cout << "sizeof(variant): " << sizeof(v) << '\n';

        v = 1.1;
        std::cout << "valueless_by_exception: " << v.valueless_by_exception() << '\n';
    }

    // ***

    {
        std::any a;
        std::cout << "sizeof(any): " << sizeof(a) << '\n';
        a = 10;
        std::cout << "type(a): " << a.type().hash_code() << '\n';
        a = "string";
        std::cout << "type(a): " << a.type().hash_code() << '\n';
    }

    return 0;
}
