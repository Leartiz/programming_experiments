#ifndef MYTH_H
#define MYTH_H

#include <QThread>
#include <QObject>

class MyTh : public QThread
{
    Q_OBJECT

public:
    explicit MyTh(
            QObject *parent = nullptr);

signals:
    void progress(int);

    // QThread interface
protected:
    void run() override;
};

#endif // MYTH_H
