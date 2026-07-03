#include <QDebug>
#include <QMetaObject>

#include <QEvent>
#include <QObject>

#include "myeventfilter.h"

MyEventFilter::MyEventFilter(QObject *parent)
    : QObject{parent} {}

bool MyEventFilter::eventFilter(QObject *watched, QEvent *event)
{
    if (event->type() == QEvent::Close) {
        auto mo = watched->metaObject();
        qDebug() << "watched type:" << mo->className();

        // ***

        if (mo->className() == QString{ "QLabel" })
            emit bb();
    }

    return QObject::eventFilter(
        watched, event);
}
