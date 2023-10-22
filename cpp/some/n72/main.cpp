#include <string>
#include <sstream>
#include <iostream>

using namespace std;

int main()
{
    std::stringstream sin{ "abcde" };
    double d{ 0 };
    sin >> d;

    std::cout << "good: " << sin.good() << "\n";
    std::cout << "fail: " <<  sin.fail() << "\n";
    std::cout << "eof: " << sin.good() << "\n";
    std::cout << "bad: " << sin.bad() << "\n";

    std::cout << "d: " << d << "\n";

    std::cout << "\n***\n\n";

    std::string buff;
    sin >> buff;

    std::cout << "good: " << sin.good() << "\n";
    std::cout << "fail: " <<  sin.fail() << "\n";
    std::cout << "eof: " << sin.good() << "\n";
    std::cout << "bad: " << sin.bad() << "\n";

    std::cout << "buff: " << '\'' << buff << '\'' << "\n";

    return 0;
}
