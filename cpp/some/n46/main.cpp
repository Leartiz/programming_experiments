#include <QList>
#include <QVector>

#include <QObject>

int main()
{
    QObject obj;

    QList<QObject> objList;
    objList.push_back(obj);

    QVector<QObject> objVec;
    objVec.push_back(obj);

    return 0;
}
