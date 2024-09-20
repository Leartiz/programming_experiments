#include <chrono>
#include <iostream>

int main()
{
    using namespace std::chrono_literals;
    constexpr auto d1{ 250ms };
    constexpr std::chrono::milliseconds d2{ 1s };

    std::cout << d1 << " = " << d1.count() << " milliseconds\n"
              << d2 << " = " << d2.count() << " milliseconds\n";
}
