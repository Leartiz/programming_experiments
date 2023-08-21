#include <QApplication>

#include "myprogressbar.h"
#include "mythread.h"

int main(int argc, char *argv[])
{
    QApplication
            a(argc, argv);

    MyProgressBar mpb;
    MyThread mth{
        &mpb
    };

    QObject::connect(&mth, &QThread::finished,
                     qApp, &QApplication::quit);

    mpb.show();
    mth.start();

    return a.exec();
}
