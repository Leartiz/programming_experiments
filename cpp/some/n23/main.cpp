#include <stdexcept>

#include <QGuiApplication>
#include <QQmlApplicationEngine>
#include <QDebug>

#include "myclass.h"

int main(int argc, char *argv[])
{
    qmlRegisterType<MyClass>("MyClass", 1, 0, "MyClass");

    // ***

#if QT_VERSION < QT_VERSION_CHECK(6, 0, 0)
    QCoreApplication::setAttribute(Qt::AA_EnableHighDpiScaling);
#endif

    // ***

    int res{ 0 };
    try {

        QGuiApplication app(argc, argv);

        QQmlApplicationEngine engine;
        const QUrl url(QStringLiteral("qrc:/main.qml"));
        QObject::connect(&engine, &QQmlApplicationEngine::objectCreated,
                         &app, [url](const QObject *obj, const QUrl &objUrl) {
            if (!obj && url == objUrl)
                QCoreApplication::exit(-1);

        }, Qt::QueuedConnection);
        engine.load(url);

        res = app.exec();
    }
    catch (const std::runtime_error& err) {
        qDebug() << "main, err: " << err.what();
    }

    return res;
}
