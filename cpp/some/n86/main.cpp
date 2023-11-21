#include <QString>

#include <QFuture>
#include <QFutureWatcher>
#include <QPromise>

#include <QDebug>
#include <QCoreApplication>
#include <QtConcurrent/QtConcurrent>

int main()
{
    auto future = QtConcurrent::run([](QPromise<QString>& promise) -> void {
        forever {

            qDebug() << ".";
            QThread::msleep(250);

            if (promise.isCanceled()) {
                return;
            }

            // ...
        }
    });

    // ***

    QThread::msleep(1500);
    future.cancel();

    future.waitForFinished();
}
