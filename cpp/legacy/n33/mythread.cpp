#include <QCoreApplication>

#include "mythread.h"
#include "progressevent.h"
#include "myprogressbar.h"

MyThread::MyThread(
        MyProgressBar *pbar,
        QObject *parent)
    : QThread(parent)
    , m_pbar{ pbar }
{}

void MyThread::run()
{
    for (int i = 0; i <= 100; ++i) {
        msleep(50);

        ProgressEvent* pe =
                new ProgressEvent;
        pe->setValue(i);

        QCoreApplication::postEvent(
                    m_pbar, pe);
    }
}
