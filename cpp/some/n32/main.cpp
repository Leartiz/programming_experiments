#include <QApplication>
#include <QProgressBar>

#include "myth.h"

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);

    QProgressBar pb;
    pb.setMaximum(100);

    // ***

    MyTh thread;
    QObject::connect(&thread, &MyTh::progress,
                     &pb, &QProgressBar::setValue);
    QObject::connect(&thread, &MyTh::finished,
                     &a, &QApplication::quit);

    // ***

    pb.show();

    thread.start();
    return a.exec();
}
