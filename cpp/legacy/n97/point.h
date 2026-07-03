#ifndef POINT_H
#define POINT_H

#include <string>

struct Point;
struct NamedPoint;

// NamedPoint --|> Point

// -----------------------------------------------------------------------

Point* newPoint(double x, double y);
NamedPoint* newNamedPoint(double x, double y, const std::string& name);

void deletePoint(Point* point);
void deleteNamedPoint(NamedPoint* point); // why?

double getPointX(Point* point);
double getPointY(Point* point);
const std::string& getName(NamedPoint* point);

void setPointX(Point* point, double x);
void setPointY(Point* point, double y);

void printPoint(Point* point);
void printNamedPoint(NamedPoint* point); // next example used function pointer!

double distanceBetweenPoints(Point* pointA,
                             Point* pointB);

#endif // POINT_H
