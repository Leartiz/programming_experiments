#include <stdexcept>

#include <QApplication>
#include <QThread>
#include <QDebug>

#include "mainwindow.h"

int main(int argc, char *argv[])
{
    int returnCode{ 0 };
    try {
        qDebug() << "main thread id:"
                 << QThread::currentThreadId();

        // ***

        QApplication a(argc, argv);
        MainWindow w; w.show();
        returnCode = a.exec();

    } catch (std::runtime_error& err) {
        qDebug() << "err:" << err.what();
    }
    return returnCode;
}
