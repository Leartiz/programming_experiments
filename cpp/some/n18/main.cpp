#include <string>
#include <iostream>

using namespace std;

void printlnStr(const std::string& str) {
    for (size_t i = 0; i < str.length(); ++i) {
        std::cout << static_cast<int>(str[i]) << " ";
    } std::cout << std::endl;
}

int main()
{
    std::string str = {
        "my str"
    };

    std::cout << str.capacity() << std::endl;
    std::cout << str.size() << std::endl;

    std::cout << std::endl;

    // ***

    str.reserve(100);
    std::cout << str << "\n";
    printlnStr(str);

    std::cout << std::endl;

    // ***

    str.resize(25);
    std::cout << str << "\n";
    printlnStr(str);

    std::cout << std::endl;

    // ***

    str.resize(2);
    std::cout << str << "\n";
    printlnStr(str);

    std::cout << std::endl;

    // ***

    std::cout << str.capacity() << std::endl;
    std::cout << str.size() << std::endl;

    return 0;
}
