#include <QDebug>
#include <QMetaEnum>

#include <QTime>
#include <QString>

#include <QThread>
#include <QWebSocket>

#include "worker.h"

namespace
{

void messageToDebug(const QString& message)
{
    qDebug() << QTime::currentTime().toString("hh:mm:ss.z")
             << message << "cur thread:" << QThread::currentThreadId();
}

} // <anonymous>

// -----------------------------------------------------------------------

Worker::Worker(int index, QObject *parent)
    : QObject{ parent }, m_index{ index } {}

void Worker::almostStart()
{
    messageToDebug("almost start");
}

void Worker::almostFinish()
{
    messageToDebug("almost finish");
    for (const auto& sk : m_wsocks)
        sk->abort();
    deleteLater();
}

void Worker::pushSock(int index, QWebSocket* sock)
{
    if (index != m_index) return;
    messageToDebug("push sock");

    sock->moveToThread(thread());
    sock->setParent(this);

    m_wsocks.push_back(sock);
    {
        connect(sock, &QWebSocket::textMessageReceived,
                this, &Worker::onTextMessageReceived_sock);
        connect(sock, &QWebSocket::disconnected,
                this, &Worker::onDisconnected_sock);
        connect(sock, static_cast<void(QWebSocket::*)(
                          QAbstractSocket::SocketError)>(
                            &QWebSocket::errorOccurred),
                this, &Worker::onErrorOccurred_sock);
    }
}

void Worker::onTextMessageReceived_sock(const QString& message)
{
    messageToDebug("text message received");
    const auto wsk = qobject_cast<QWebSocket*>(sender());

    QThread::msleep(250);

    emit messageProcessed(message);
    wsk->sendTextMessage("+"); // results in an error: --------------------------
    //                                                                          \
    //                                                                          \
    // Socket notifiers cannot be enabled or disabled from another thread... <---
}

// handlers don't work...

void Worker::onDisconnected_sock()
{
    messageToDebug("disconnected");
    const auto wsk = qobject_cast<QWebSocket*>(sender());
    m_wsocks.removeOne(wsk);
}

void Worker::onErrorOccurred_sock(QAbstractSocket::SocketError err)
{
    const auto metaEnum = QMetaEnum::fromType<QAbstractSocket::SocketError>();
    messageToDebug("error occurred: " + QString{ metaEnum.key(err) });
}
