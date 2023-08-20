#include <QDebug>

#include "mythread.h"

MyThread::MyThread(QObject *parent)
    : QThread{ parent }
{
    // to error:
    // "QObject::startTimer: Timers cannot be started
    //     from another thread"

    // m_timer = new QTimer{};
}

void MyThread::onTimeout_timer()
{
    qDebug() << "MyThread, onTimeout_timer, cur thread id:"
             << QThread::currentThreadId();

    emit counterChanged(
        m_counter++);
}

void MyThread::run()
{
    qDebug() << "MyThread, run, cur thread id:"
             << QThread::currentThreadId() << "\n";

    m_timer = new QTimer{};
    m_timer->setInterval(500);
    m_timer->setSingleShot(false);
    {
        connect(m_timer, &QTimer::timeout,
                this, &MyThread::onTimeout_timer);
    }

    // ***

    m_timer->start();

    // ***

    exit(
        exec());
}
