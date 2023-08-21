#include "myth.h"

MyTh::MyTh(QObject *parent)
    : QThread(parent) {}

void MyTh::run()
{
    for (int i = 0; i <= 100; ++i) {
        msleep(10); emit progress(i);
    }
}
