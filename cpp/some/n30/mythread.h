#ifndef MYTHREAD_H
#define MYTHREAD_H

#include <QThread>
#include <QObject>
#include <QTimer>

class MyThread : public QThread
{
    Q_OBJECT

public:
    explicit MyThread(
        QObject *parent = nullptr);

signals:
    void counterChanged(int);

public slots:
    void onTimeout_timer();

    // QThread interface
protected:
    void run();

private:
    QTimer* m_timer;
    int m_counter{ 0 };
};

#endif // MYTHREAD_H
