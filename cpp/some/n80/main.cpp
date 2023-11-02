#include <QDebug>
#include <QThreadPool>

int main()
{
    auto thPool = QThreadPool::globalInstance();
    qDebug() << thPool->activeThreadCount();
    qDebug() << thPool->maxThreadCount();

    // ***




    return 0;
}
