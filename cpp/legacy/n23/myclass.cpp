#include <stdexcept>

#include <QDebug>

#include "myclass.h"

MyClass::MyClass(QObject* parent)
    : QObject{ parent } {}

void MyClass::initWithException()
{
    qDebug() << "MyClass, try init with exception";

    // ...

    throw std::runtime_error("something broke");
}

void MyClass::init()
{
    qDebug() << "MyClass, try init";

    // ...

    emit initFailed("something broke");
}
