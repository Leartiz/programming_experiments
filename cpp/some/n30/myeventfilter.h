#ifndef MYEVENTFILTER_H
#define MYEVENTFILTER_H

#include <QObject>

class MyEventFilter : public QObject
{
    Q_OBJECT
public:
    explicit MyEventFilter(QObject *parent = nullptr);

    // QObject interface
public:
    bool eventFilter(QObject *watched, QEvent *event);

signals:
    void bb();
};

#endif // MYEVENTFILTER_H
