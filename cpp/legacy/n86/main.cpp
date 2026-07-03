#include <QString>

#include <QFuture>
#include <QPromise>

#include <QDebug>
#include <QCoreApplication>
#include <QtConcurrent/QtConcurrent>

int main()
{
    auto future = QtConcurrent::run([](QPromise<QString>& promise) -> void {
        forever {

            qDebug() << ".";
            QThread::msleep(100);

            if (promise.isCanceled()) {
                return;
            }

            // ...
        }
    });

    // ***

    QThread::msleep(500);
    future.cancel();

    future.waitForFinished();
}
