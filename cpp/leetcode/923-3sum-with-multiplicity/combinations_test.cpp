// go test style: build and run combinations_test

#include <cassert>
#include <iostream>

#include "combinations.hpp"

int main()
{
    assert(combinations(5, 0) == 1);
    assert(combinations(5, 1) == 5);
    assert(combinations(5, 2) == 10);
    assert(combinations(5, 3) == 10); // symmetry: C(5,3) == C(5,2)
    assert(combinations(5, 5) == 1);

    assert(combinations(4, 2) == 6);
    assert(combinations(4, 3) == 4); // 923: pick 3 of four equal values

    assert(combinations(3, 5) == 0); // k > n
    assert(combinations(5, -1) == 0);

    std::cout << "combinations_test: ok\n";
}
