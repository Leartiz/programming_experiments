#include <QDebug>

#include "myclass.h"

MyClass::MyClass(QObject *parent)
    : QObject(parent)
{
    qDebug() << "MyClass[" + objectName() + "]::ctor()" ;
}

MyClass::~MyClass()
{
    qDebug() << "MyClass[" + objectName() + "]::dtor()" ;
}
