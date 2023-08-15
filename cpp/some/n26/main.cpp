#include <QDebug>
#include <QObject>
#include <QString>
#include <QTimer>

#include <QCoreApplication>

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);

    const auto metaCoreApp = a.metaObject();
    qDebug() << metaCoreApp->className();

    // ***

    const char* objClassName = "QObject";
    if (a.inherits(objClassName)) {
        qDebug() << "a is" << objClassName;
    }

    // ***

    const char* strClassName = "QString";
    if (!a.inherits(strClassName)) {
        qDebug() << "a is not" << strClassName;
    }

    // ***

    qDebug() << "\n";

    QTimer::singleShot(500, &a,
                       &QCoreApplication::quit);

    return a.exec();
}
