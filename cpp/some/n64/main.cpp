#include <QDebug>

#include <QCoreApplication>
#include <QNetworkInterface>

QString getMacAddress()
{
    QString text;
    foreach(QNetworkInterface interface, QNetworkInterface::allInterfaces()) {
        text += "Interface:"+interface.hardwareAddress()+"\n";
    }
    return text;
}

int main()
{
    qDebug() << getMacAddress();
    return 0;
}
