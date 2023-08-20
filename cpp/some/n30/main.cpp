#include <iostream>

#include <QObject>
#include <QApplication>

#include <QLabel>
#include <QPushButton>

#include "mythread.h"
#include "myeventfilter.h"

using namespace std;

int main(int argc, char** argv)
{
    QApplication app{ argc, argv };

    // ***

    qDebug() << "main, cur thread id:"
             << QThread::currentThreadId();

    MyThread mth;

    QLabel label; // main thread!
    label.resize(250, 250);
    label.setText("0x00000000");

    // ***

    MyEventFilter mef;
    label.installEventFilter(&mef); // label is widget!

    // ***

    QObject::connect(&mth, &MyThread::counterChanged, &label,
                     [&label](int value) -> void {

        qDebug() << "QLabel, counterChanged -> lambda, "
                    "cur thread id:"
                 << QThread::currentThreadId();

        label.setNum(value);
    });

    QObject::connect(&mef, &MyEventFilter::bb, &mef,
                     [&mth]() -> void {

        qDebug() << "MyEventFilter, bb -> lambda";

        mth.quit();
    });

    mth.moveToThread(&mth); // does not move the context?

    // ***

    label.show();
    mth.start();

    return app.exec();
}
