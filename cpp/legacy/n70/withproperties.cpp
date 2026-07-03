#include <QDebug>

#include <QMetaObject>
#include <QMetaProperty>

#include "withproperties.h"

WithProperties::WithProperties(QObject* parent)
    : QObject{ parent } {}

WithProperties::~WithProperties() {};

void WithProperties::printProps() {
    const auto mo = metaObject();

    const auto pCount = mo->propertyCount();
    const auto pOffset = mo->propertyOffset();

    qDebug() << "propertyCount:" << pCount;
    qDebug() << "propertyOffset:" << pOffset;

    // ***

    qDebug() << "static props {";
    for (int i = pOffset; i <= pOffset + mo->propertyCount(); ++i) {
        qDebug() << mo->property(i).name();
    }
    qDebug() << "}";

    // ***

    qDebug() << "dynamic props {";
    const auto dynProps = dynamicPropertyNames();
    for (int i = 0; i < dynamicPropertyNames().size(); ++i) {
        qDebug() << dynProps[i];
    }
    qDebug() << "}";
}
