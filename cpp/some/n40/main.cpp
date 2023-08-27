#include <QDebug>
#include <QCoreApplication>

#include <QtConcurrent/QtConcurrent>
#include <QSharedPointer>
#include <QTimer>

int main()
{
    {
        QScopedPointer<int> ptr(
            new int(10));
    }

    // ***

    QSharedPointer<QPromise<int>>
        sharedPromise(
            new QPromise<int>());

    QFuture<int> future =
        sharedPromise->future();

    // ***

    sharedPromise->start();

    // placement?
    QScopedPointer<QThread> threads[] = {
        QScopedPointer<QThread>(QThread::create([] (auto sharedPromise) {
            sharedPromise->addResult(0, 0);
        }, sharedPromise)),
        QScopedPointer<QThread>(QThread::create([] (auto sharedPromise) {
            sharedPromise->addResult(-1, 1);
        }, sharedPromise)),
        QScopedPointer<QThread>(QThread::create([] (auto sharedPromise) {
            sharedPromise->addResult(-2, 2);
        }, sharedPromise)),
        // ...
    };

    for (auto& t : threads)
        t->start();

    // ***

    // without finish, can already pick up result?
    for (int i = 0; i < 3; ++i) {
        qDebug() << "res #" << i << ":"
                 << future.resultAt(i);
    }

    // ***

    sharedPromise->finish();
}
