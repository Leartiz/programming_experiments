#include <QDebug>
#include <QMetaObject>
#include <QMetaProperty>

#include "myclass.h"

int main()
{
    MyClass mc{ nullptr };
    qDebug() << mc.isReadOnly();

    mc.setProperty("readOnly", true);
    qDebug() << mc.property("readOnly").toBool();

    // ***

    const QMetaObject* pmo = mc.metaObject();
    for (int i = 0; i < pmo->propertyCount(); ++i) {
        const QMetaProperty mp = pmo->property(i);
        qDebug() << "Property #" << i;
        qDebug() << "Type:" << mp.typeName();
        qDebug() << "Name:" << mp.name();
        qDebug() << "Val:" << mc.property(mp.name());
    }

    return 0;
}
