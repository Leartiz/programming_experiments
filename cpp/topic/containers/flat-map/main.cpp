#include <boost/container/flat_map.hpp>
#include <iostream>
#include <string>

int main() {
    boost::container::flat_map<std::string, int> m{
        {"banana", 2},
        {"apple", 1},
    };
    m["cherry"] = 3;

    for (const auto& [k, v] : m) {
        std::cout << k << " -> " << v << '\n';
    }
}
