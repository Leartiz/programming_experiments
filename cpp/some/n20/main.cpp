#include <string>
#include <vector>

#include <iostream>
#include <stdexcept>

using namespace std;

int main()
{
    std::string str;
    std::cout << str.capacity();

    str.reserve(100);
    std::cout << str.capacity();

    // memory must be initialized (?)
    try {
        str.at(99) = 'a';
    }
    catch (const std::exception& e) {
        std::cout << e.what() << "\n";
    }

    // ***

    str.resize(100); // default constructor (?)
    std::cout << str.capacity() << std::endl;
    std::cout << str.size() << std::endl;

    return 0;
}
