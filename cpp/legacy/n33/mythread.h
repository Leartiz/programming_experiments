#ifndef MYTHREAD_H
#define MYTHREAD_H

#include <QThread>
#include <QObject>

class MyProgressBar;

class MyThread : public QThread
{
    Q_OBJECT
public:
    MyThread(
            MyProgressBar *pbar,
            QObject *parent = Q_NULLPTR);

    // QThread interface
protected:
    void run() override;

private:

    // receiver!
    MyProgressBar *m_pbar{
        nullptr
    };
};

#endif // MYTHREAD_H
