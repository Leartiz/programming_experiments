#include <QDebug>
#include <QJsonObject>
#include <QJsonDocument>

#include <QList>
#include <QListIterator>
#include <QMutableListIterator>
#include <QString>

struct Point final
{
public:
    Point(const int x, const int y)
        : x{ x }, y{ y } {}

public:
    Point(const Point& p)
        : x{ p.x }, y{ p.y }  {
        qDebug() << "Point, copy-ctor";
    }
    Point& operator=(const Point& p) {
        qDebug() << "Point, copy-operator=";
        if (&p != this) {
            this->x = p.x;
            this->y = p.y;
        }
        return *this;
    }

public:
    Point(Point&& p)
        : x{ p.x }, y{ p.y }  {
        qDebug() << "Point, move-ctor";
    }
    Point& operator=(Point&& p) {
        qDebug() << "Point, move-operator=";
        if (&p != this) {
            this->x = p.x;
            this->y = p.y;
        }
        return *this;
    }

public:
    QString toString() const {
        QJsonObject jsonObj;
        jsonObj["x"] = x;
        jsonObj["y"] = y;
        return QJsonDocument{ jsonObj }
            .toJson(QJsonDocument::Compact);
    }

public:
    int x{ 0 }, y{ 0 };
};

int main()
{
    QList<QString> strs;
    strs << "1" << "2" << "3";

    // java
    {
        QListIterator<QString> it{ strs };
        it.findNext("2");
        qDebug() << "unit:" << it.previous() << "\n";

        // ***

        it.toFront();
        while (it.hasNext()) {
            qDebug() << "unit:" << it.next();
        }
    }
    qDebug() << "---";
    {
        QListIterator<QString> it{ strs };
        it.toBack();
        while (it.hasPrevious()) {
            qDebug() << "unit:" << it.previous();
        }
    }
    {

    }

    qDebug() << "---";

    // stl
    {
        for (auto it = strs.begin(); it != strs.end(); ++it) {
            qDebug() << "unit:" << *it;
        }
    }

    qDebug() << "---";

    // foreach
    {
        QList<Point> points;
        points.reserve(3);
        points.push_back(Point{ 1, 1 });
        points.push_back(Point{ 2, 2 });
        points.push_back(Point{ 3, 3 });

        qDebug() << "***";

        foreach (const Point& one, points) {
            qDebug() << "unit:"
                     << one.toString();
        }

        qDebug() << "***";

        foreach (Point one, points) {
            one.x = 10;
            qDebug() << "unit:"
                     << one.toString();
        }
    }

    return 0;
}
