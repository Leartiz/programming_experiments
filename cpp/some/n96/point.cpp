#include <stdio.h>

#include "point.h"

struct Point
{
    double x;
    double y;
};

// ctor and dtor
// -----------------------------------------------------------------------

Point* newPoint(double x, double y)
{
    return new Point{
        .x = x,
        .y = y
    };
}

void deletePoint(Point* point)
{
    if (point != nullptr) {
        delete point;
    }
}

// -----------------------------------------------------------------------

double getPointX(Point* point)
{
    return point->x;
}

double getPointY(Point* point)
{
    return point->y;
}

void setPointX(Point* point, double x)
{
    point->x = x;
}

void setPointY(Point* point, double y)
{
    point->y = y;
}

// -----------------------------------------------------------------------

void movePoint(Point* point,
               double x, double y)
{
    point->x += x;
    point->y += y;
}

void printPoint(Point* point)
{
    printf("Point: { x: %f, y: %f }",
           point->x, point->y);
}
