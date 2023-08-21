#include <QLCDNumber>
#include <QApplication>

#include <QThread>

#include "myworker.h"

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    QLCDNumber lcd;

    QThread th;
    MyWorker wk;

    QObject::connect(&wk, &MyWorker::valueChanged,
                     &lcd, [&lcd](int value) -> void {
        lcd.display(value);
    });

    lcd.display("");
    lcd.resize(200, 100);
    lcd.show();

    wk.moveToThread(&th);

    QObject::connect(&th, &QThread::started,
                     &wk, &MyWorker::doWork);

    QObject::connect(&th, &QThread::finished,
                     &wk, &MyWorker::doStop);

    QObject::connect(&wk, &MyWorker::made,
                     &a, &QApplication::quit);

    th.start();

    // ***

    const int result =
            a.exec();

    th.quit();
    th.wait();

    return result;
}
