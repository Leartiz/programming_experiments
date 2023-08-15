#include <QDebug>
#include <QObject>
#include <QVector>

#include <QTimer>
#include <QCoreApplication>

#include "myclass.h"

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);
    a.setObjectName("core app"); // else <empty>

    MyClass* obj0 = new MyClass(&a);
    MyClass* obj1 = new MyClass(obj0);
    MyClass* obj2 = new MyClass(obj1);
    MyClass* obj3 = new MyClass(obj0);
    MyClass* obj4 = new MyClass(obj2);

    // ***

    QVector<QObject*> objs;
    objs.push_back(obj0);
    objs.push_back(obj1);
    objs.push_back(obj2);
    objs.push_back(obj3);
    objs.push_back(obj4);

    for (int i = 0; i < objs.size(); ++i) {
        qDebug() << objs[i]->objectName();
    }
    qDebug() << "\n";

    // ***

    obj0->setObjectName("obj0");
    obj1->setObjectName("obj1, child obj0");
    obj2->setObjectName("obj2, child obj1");
    obj3->setObjectName("obj3, child obj0");
    obj4->setObjectName("obj4, child obj2");

    for (int i = 0; i < objs.size(); ++i) {
        qDebug() << objs[i]->objectName();
    }
    qDebug() << "\n";

    // ***

    for (QObject* obj = obj4; obj; obj = obj->parent()) {
        qDebug() << obj->objectName();
    }
    qDebug() << "\n";

    // ***

    obj0->dumpObjectTree();
    delete obj0;

    // ***

    QTimer::singleShot(500, &a, &QCoreApplication::quit);

    qDebug() << int('\n');
    qDebug() << int('\r');

    return a.exec();
}
