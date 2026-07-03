#include <iostream>

#include "point.h"

int main()
{
    {
        auto *p1 = newPoint(0, 100);
        auto *p2 = newNamedPoint(0, 10, "Test");

        std::cout << "dist: " << distanceBetweenPoints(
            p1, reinterpret_cast<Point*>(p2)) << "\n";

        deletePoint(p1);
        deleteNamedPoint(p2);
    }

    {
        auto *p = newNamedPoint(0, 10, "Test");
        setPointX(reinterpret_cast<Point*>(p), 155);

        printNamedPoint(p);
    }
}
