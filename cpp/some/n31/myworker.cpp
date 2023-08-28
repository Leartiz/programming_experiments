#include "myworker.h"

MyWorker::MyWorker(QObject *parent)
    : QObject(parent)
{
    m_timer = new QTimer{ this }; // allows to move to another thread!

    m_timer->setSingleShot(false);
    m_timer->setInterval(250);
    {
        connect(m_timer, &QTimer::timeout,
                this, &MyWorker::onTimeout_timer);

        connect(this, &MyWorker::made,
                m_timer, &QTimer::stop);
    }
}

void MyWorker::doWork()
{
    m_timer->start();
}

void MyWorker::doStop()
{
    m_timer->stop();
}

void MyWorker::onTimeout_timer()
{
    if (!m_value) {
        emit made();
    }
    emit valueChanged(
                m_value--);
}

