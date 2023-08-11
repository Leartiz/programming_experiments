#ifndef MYCLASS_H
#define MYCLASS_H

#include <QDebug>
#include <QObject>

class MyClass final
        : public QObject
{
     Q_OBJECT

public slots:
    void slotPb() {
        qDebug() << "slotPb";
    }

protected slots:
    void slotPo() {
        qDebug() << "slotPo";
    }

private slots:
    void slotPi() {
        qDebug() << "slotPi";
    }
};


#endif // MYCLASS_H
