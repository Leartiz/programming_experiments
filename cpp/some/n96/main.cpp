#include <iostream>

#include "point.h"

int main()
{
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
    }

    // ***

    {
        auto *p1 = newPoint(0, 100);
        auto *p2 = newPoint(0, 10);

        std::cout << "dist: " << distanceBetweenPoints(p1, p2);

        deletePoint(p1);
        deletePoint(p2);
    }

    return 0;
}
