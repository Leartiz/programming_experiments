#ifndef PROGRESSEVENT_H
#define PROGRESSEVENT_H

#include <QEvent>
#include <QObject>

class ProgressEvent :
        public QEvent
{
public:
    enum {
        Progress = User + 1,
    };

public:
    ProgressEvent();
    void setValue(int);
    int value() const;

private:
    int m_value{ 0 };
};

#endif // PROGRESSEVENT_H
