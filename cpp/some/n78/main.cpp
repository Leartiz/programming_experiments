#include <iostream>
#include <variant>

int main()
{
    std::variant<int> v;
    v = 100;

    std::cout << std::get<int>(v);
    return 0;
}
