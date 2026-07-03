#include <QThread>
#include <QCoreApplication>
#include <QDateTime>
#include <QTimeZone>

#include <QDebug>

class SvrController
{

};


int main(int argc, char *argv[])
{
    QDateTime UTC(QDateTime::fromSecsSinceEpoch(1698266256, QTimeZone::UTC));
    QDateTime local(UTC.toTimeZone(QTimeZone::systemTimeZone()));

    qDebug() << UTC.toString();
    qDebug() << local.toString();

    return 0;
}
