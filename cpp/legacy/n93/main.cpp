#include <QDebug>

#include <QFuture>
#include <QtConcurrent/QtConcurrent>

using namespace std;

int main()
{
    qDebug() << QThread::currentThreadId();

    auto future = QtConcurrent::run([]()-> int{
        qDebug() << QThread::currentThreadId();
        return std::rand() % 123;
    }).then([](const int value) -> QString {
        qDebug() << QThread::currentThreadId();
        return QString::number(value);
    }).then([](const QString& value) -> QByteArray {
        qDebug() << QThread::currentThreadId();
        return value.toUtf8().toHex();
    });

    qDebug() << future.result();
}
