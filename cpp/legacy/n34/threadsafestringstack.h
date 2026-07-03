#ifndef THREADSAFESTRINGSTACK_H
#define THREADSAFESTRINGSTACK_H

#include <QMutex>
#include <QStack>
#include <QString>
#include <QStringList>

class ThreadSafeStringStack
{
public:
    void push(const QString&);
    QString pop();

public:
    QStringList list();

private:
    QMutex m_mutex;
    QStack<QString> m_strs;
};

#endif // THREADSAFESTRINGSTACK_H
