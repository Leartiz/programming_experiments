#include <cmath>

#include "point.h"

struct Point
{
    double x;
    double y;
};

struct NamedPoint
{
    double x;
    double y;

    std::string name;
};

// -----------------------------------------------------------------------

Point* newPoint(double x, double y)
{
    return new Point{
        .x = x,
        .y = y,
    };
}

NamedPoint* newNamedPoint(double x, double y, const std::string& name)
{
    return new NamedPoint{
        .x = x,
        .y = y,
        .name = name,
    };
}

void deletePoint(Point* point)
{
    delete point;
}

void deleteNamedPoint(NamedPoint* point)
{
    delete point;
}

double distanceBetweenPoints(Point* pointA,
                             Point* pointB)
{
    const auto diffX = pointA->x - pointB->x;
    const auto diffY = pointA->y - pointB->y;

    return std::sqrt(diffX*diffX) +
           std::sqrt(diffY*diffY);
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

const std::string& getName(NamedPoint* point)
{
    return point->name;
}

// -----------------------------------------------------------------------

void setPointX(Point* point, double x)
{
    point->x = x;
}

void setPointY(Point* point, double y)
{
    point->y = y;
}

// -----------------------------------------------------------------------

void printPoint(Point* point)
{
    printf("Point: { x: %f, y: %f }",
           point->x, point->y);
}

void printNamedPoint(NamedPoint* point)
{
    printf("Point: { x: %f, y: %f, name: %s }",
           point->x, point->y, point->name.c_str());
}
