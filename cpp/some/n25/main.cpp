#include <QTimer>
#include <QObject>

#include <QRegularExpression>

#include <QCoreApplication>

#include "myclass.h"

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);
    a.setObjectName("core app");

    MyClass* obj0 = new MyClass(&a);
    obj0->setObjectName("obj0, my class");

    QObject* obj1 = new QObject(obj0);
    obj1->setObjectName("obj1");

    QObject* obj2 = new QObject(obj0);
    obj2->setObjectName("obj2");

    QObject* obj3 = new QObject(obj1);
    obj3->setObjectName("obj3"); // <---------------
                                        //         |
    QObject* obj31 = new QObject(obj1); //         |
    obj31->setObjectName("obj3"); // same ----------

    // spec case
    QObject* specObj0 = new QObject(obj2);
    specObj0->setObjectName("specObj0");

    qDebug() << "\n";

    // ***

    obj0->dumpObjectInfo();
    qDebug() << "\n";
    obj0->dumpObjectTree();

    qDebug() << "\n";

    // ***

    obj1->dumpObjectInfo();
    qDebug() << "\n";
    obj1->dumpObjectTree();

    qDebug() << "\n";

    // ***

    auto childs = obj0->findChildren<QObject*>();
    for (int i = 0; i < childs.size(); ++i) {
        qDebug() << childs[i]->objectName();
    }

    qDebug() << "\n";

    // ***

    qDebug() << "without spec object!";
    childs = obj0->findChildren<QObject*>(
                QRegularExpression(R"(obj\d+)"));
    for (int i = 0; i < childs.size(); ++i) {
        qDebug() << childs[i]->objectName();
    }

    qDebug() << "\n";

    // ***

    auto childObj3 = obj0->findChild<QObject*>("obj3");
    if (childObj3 == obj3) {
        qDebug() << "found obj3";
    }
    if (childObj3 == obj31) {
        qDebug() << "found obj31";
    }

    qDebug() << "\n";

    // ***

    QTimer::singleShot(500, &a,
                       &QCoreApplication::quit);

    return a.exec();
}
