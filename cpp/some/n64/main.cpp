#include <QDebug>
#include <QTextStream>

#include <QCoreApplication>
#include <QNetworkInterface>

QString getMacAddress()
{
    QString text;
    QTextStream textStream{ &text };

    foreach(QNetworkInterface interface, QNetworkInterface::allInterfaces())
        textStream << "interface: " + interface.hardwareAddress() << "\n";
    return text;
}

int main()
{
    const auto some = getMacAddress().split("\n", Qt::SplitBehaviorFlags::SkipEmptyParts);
    for (int i = 0; i < some.size(); ++i)
        qDebug() << some[i];
    return 0;
}
