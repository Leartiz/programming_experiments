#include <QDebug>
#include <QObject>

#include <QCoreApplication>

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);

    const auto metaCoreApp = a.metaObject();
    qDebug() << metaCoreApp->className();

    return a.exec();
}
