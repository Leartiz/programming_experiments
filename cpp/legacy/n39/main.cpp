#include <functional>

#include <QDebug>
#include <QObject>
#include <QTimer>

#include <QCoreApplication>
#include <QtConcurrent/QtConcurrent>

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);

    // ***

    QPromise<int> p;

    QFuture<int> f1 = p.future();
    f1.then([](int val) {
        qDebug("first");
        qDebug() << val;
    });

    QFuture<int> f2 = p.future();
    f2.then([](int val) {
        qDebug("second");
        qDebug() << val;
    });

    auto closeAppByNeed = [&](QObject* sender) {
        const auto mo = sender->metaObject();
        qDebug() << "closeAppByNeed, class name sender:"
                 << mo->className();

        if (dynamic_cast<QTimer*>(sender)) {
            qDebug() << "close app by timer";
            qApp->quit();
        }

        if (f1.isFinished() || f2.isFinished()) {
            qDebug() << "close app by future";
            qApp->quit();
        }
    };

    // ***

    QFutureWatcher<int> fw1;
    fw1.setFuture(f1);

    QFutureWatcher<int> fw2;
    fw1.setFuture(f2);

    QObject::connect(&fw1, &QFutureWatcher<int>::finished,
                     qApp, std::bind(closeAppByNeed, &fw1));

    QObject::connect(&fw2, &QFutureWatcher<int>::finished,
                     qApp, std::bind(closeAppByNeed, &fw2));

    // ***

    QTimer timer;
    timer.setInterval(1000);
    timer.setSingleShot(true);
    QObject::connect(&timer, &QTimer::timeout,
                     qApp, std::bind(closeAppByNeed, &timer));

    // ***

    p.start();
    p.addResult(55);
    p.finish();

    // ***

    timer.start();
    return a.exec();
}
