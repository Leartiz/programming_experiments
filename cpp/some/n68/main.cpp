#include <QDateTime>
#include <QCoreApplication>
#include <QThread>

int main()
{    
    QDateTime dt1 = QDateTime::currentDateTime();

    QThread::sleep(3);

    QDateTime dt2 = QDateTime::currentDateTime();

    qDebug() <<
                dt2.toSecsSinceEpoch() -
                dt1.toSecsSinceEpoch();

    return 0;
}
