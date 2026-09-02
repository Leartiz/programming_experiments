#include "thread_safe_cache.hpp"

#include <iostream>
#include <string>

int main()
{
    ThreadSafeCache<std::string, int> cache;

    cache.put("x", 42);

    std::cout << cache.get("x") << '\n';
    std::cout << cache.contains("x") << '\n';
}
