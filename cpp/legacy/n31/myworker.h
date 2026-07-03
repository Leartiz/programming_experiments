#ifndef MYWORKER_H
#define MYWORKER_H

#include <QTimer>
#include <QObject>

class MyWorker : public QObject
{
    Q_OBJECT

public:
    explicit MyWorker(
            QObject *parent = nullptr);

signals:
    void made();
    void valueChanged(int);

public slots:
    void doWork();
    void doStop();

private slots:
    void onTimeout_timer();

private:
    int m_value{ 10 };
    QTimer* m_timer{ nullptr };
};

#endif // MYWORKER_H
