#include <QDebug>

#include <QMetaObject>
#include <QMetaProperty>

#include "withproperties.h"

int main()
{
    WithProperties wp;
    wp.printProps();

    // ***

    {
        wp.setProperty(qPrintable(QString{ "ум" }), QVariant::fromValue(412));
        wp.setProperty(qPrintable(QString{ "кусок" }), QVariant::fromValue(124));
        wp.setProperty(qPrintable(QString{ "владелец" }), QVariant::fromValue(125));

        wp.setProperty(qPrintable(QString{ "ratio" }), QVariant::fromValue(123));
        wp.setProperty(qPrintable(QString{ "manufacturer" }), QVariant::fromValue(124));
        wp.setProperty(qPrintable(QString{ "truth" }), QVariant::fromValue(125));
    }

    // ***

    wp.printProps();
}
