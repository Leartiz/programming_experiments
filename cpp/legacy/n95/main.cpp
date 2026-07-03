#include <iostream>
#include <cstdint>
#include <string>
#include <vector>

#include <bit>
#include <bitset>
#include <limits>

#include "someenum.h"
#include "enumarray.h"

using namespace std;

template<typename T>
void unordered_erase(std::vector<T>& v, int index)
{
    std::swap(v[index], v.back());
    v.pop_back();
}

// -----------------------------------------------------------------------

int main()
{
    const auto value = std::bit_cast<uint64_t, double>(123.123);
    std::cout << std::bitset<sizeof(value) * 8>(value) << "\n";

    // ***

    std::cout << "float eps: " << std::numeric_limits<float>::epsilon() << "\n";
    std::cout << "double eps: " << std::numeric_limits<double>::epsilon() << "\n";
    std::cout << "int eps: " << std::numeric_limits<int>::epsilon() << "\n";

    // ***

    std::cout << "enum count: " << (int)SomeEnum::Count << "\n";

    // ***

    EnumArray<SomeEnum, std::string> names{
        { SomeEnum::Grams, std::string{ "G" } },
    };
    std::cout << names[SomeEnum::Grams] << "\n";

    return 0;
}
