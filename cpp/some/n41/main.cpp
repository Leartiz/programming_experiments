#include <QDebug>
#include <QVariant>
#include <QDateTime>

int main()
{
    QVariant v;

    v = 1;
    qDebug() << v.toInt();

    v = "123";
    qDebug() << v.toString();

    v = true;
    qDebug() << v.toBool();

    v = double{ 1.1 };
    qDebug() << "double as int:"    << v.toInt();
    qDebug() << "double as double:" << v.toDouble();

    v = QDateTime::currentDateTime();
    qDebug() << "date time:" << v.toDateTime().toString("dd.MM.yyyy hh:mm:ss.z");

    return 0;
}
