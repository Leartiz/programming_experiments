#include <QDebug>
#include <QTextStream>

#include <QCoreApplication>
#include <QNetworkInterface>

QString getMacAddress()
{
    QString text;
    QTextStream textStream{ &text };

    foreach(QNetworkInterface interface, QNetworkInterface::allInterfaces())
        textStream << "interface: " + interface.hardwareAddress() << " ";
    return text;
}

int main()
{
    qDebug() << getMacAddress();
    return 0;
}
