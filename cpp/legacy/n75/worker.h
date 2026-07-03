#ifndef WORKER_H
#define WORKER_H

#include <QObject>
#include <QVector>
#include <QWebSocket>

class Worker : public QObject
{
    Q_OBJECT
public:
    explicit Worker(int index, QObject *parent = nullptr);


signals:
    void messageProcessed(const QString&);

public slots:
    void almostStart();
    void almostFinish();
    void pushSock(int, QWebSocket*);

public slots:
    void onTextMessageReceived_sock(const QString&);
    void onDisconnected_sock();
    void onErrorOccurred_sock(QAbstractSocket::SocketError);

private:
    int m_index;
    QVector<QWebSocket*> m_wsocks;
};

#endif // WORKER_H
