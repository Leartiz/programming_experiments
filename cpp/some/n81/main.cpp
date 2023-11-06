#include <iostream>
#include <stdexcept>

class A final {
public:
    A() try {
        throw std::runtime_error{"something broke"};
    }
    catch (const std::runtime_error& err) {
        std::cerr << "A ctor err:" << err.what() << std::endl;

        // implicit ---
        //            |
        //   -----------
        //   |
        //   V
        // throw;
    }
};

void foo1()
{
    throw 1;
}

void foo()
{
    foo1();
    //throw;
}

int main()
{
    try {
        const auto a = A();
        static_cast<void>(a);
    }
    catch(const std::runtime_error& err) {
        std::cerr << "main err: " << err.what() << std::endl;
    }

    std::cout << "[OK]" << std::endl;
}
