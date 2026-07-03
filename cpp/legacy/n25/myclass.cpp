#include <QDebug>

#include "myclass.h"

MyClass::MyClass(QObject *parent)
    : QObject(parent)
{
    // this
    {
        connect(this, &MyClass::innerValuesChanged,
                &MyClass::onInnerValuesChanged);
    }
}

void MyClass::onInnerValuesChanged()
{
    qDebug() << "MyClass, on inner values changed";
}
