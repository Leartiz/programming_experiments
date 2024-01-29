#ifndef POINT_H
#define POINT_H

struct Point;

Point* newPoint(double x, double y);
void deletePoint(Point* point);

double getPointX(Point* point);
double getPointY(Point* point);

void setPointX(Point* point, double x);
void setPointY(Point* point, double y);

void movePoint(Point* point,
               double x, double y);
void printPoint(Point* point);

double distanceBetweenPoints(Point* pointA,
                             Point* pointB);

#endif // POINT_H
