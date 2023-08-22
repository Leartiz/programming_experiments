#include <QMutexLocker>

#include "threadsafestringstack.h"

void ThreadSafeStringStack::push(const QString& one)
{
    QMutexLocker _(&m_mutex);
    m_strs.push(one);
}

QString ThreadSafeStringStack::pop()
{
    QMutexLocker _(&m_mutex);
    return m_strs.empty() ? QString{} : m_strs.pop();
}

QStringList ThreadSafeStringStack::list()
{
    QStringList list;
    QMutexLocker _(&m_mutex);
    for (int i = 0; i < m_strs.size(); ++i) {
        list.push_back(
                    m_strs[i]);
    }
    return list;
}
