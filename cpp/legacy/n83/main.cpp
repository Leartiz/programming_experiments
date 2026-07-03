#include <iostream>
#include <stdexcept>

void foo() try
{
    throw std::runtime_error{"something broke"};
}
catch (const std::runtime_error& err)
{
    std::cerr << "foo err: " << err.what() << std::endl;
}

void safe_foo() noexcept(bool(1 + 1))
{
    // actions...
}

int main()
{
    foo();

    safe_foo();

    std::cout << "[OK]" << std::endl;
    return 0;
}
