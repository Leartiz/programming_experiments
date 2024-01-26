#include <iostream>

#include "point.h"

int main()
{
    auto *p = newPoint(123, 123);

    std::cout << "x: " << getPointX(p) << " ";
    std::cout << "y: " << getPointY(p) << "\n";

    movePoint(p, 5, -5);
    std::cout << "x: " << getPointX(p) << " ";
    std::cout << "y: " << getPointY(p) << "\n";

    printPoint(p);
    std::cout << "\n";

    deletePoint(p);
    return 0;
}
