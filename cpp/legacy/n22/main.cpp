#include <QTimer>
#include <QObject>
#include <QCoreApplication>

#include "myclass.h"

int main(int argc, char** argv)
{
    QCoreApplication app(argc, argv);

    MyClass mc;
    mc.slotPb();

    // ***

    const auto dur = 250;

    QTimer::singleShot(dur, &mc, SLOT(slotPb())); // public
    QTimer::singleShot(dur, &mc, SLOT(slotPo())); // protected
    QTimer::singleShot(dur, &mc, SLOT(slotPi())); // private

    // ***


    QTimer::singleShot(dur, &mc, &MyClass::slotPb);
    /*

    but it won't work!

    QTimer::singleShot(dur, &mc, &MyClass::slotPo);
    QTimer::singleShot(dur, &mc, &MyClass::slotPi);

    */

    // ***

    QTimer::singleShot(dur * 3, &app, &QCoreApplication::quit);

    return app.exec();
}
