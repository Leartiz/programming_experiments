#include <QDebug>

#include <QThread>
#include <QRandomGenerator>

#include <QVector>
#include <QString>

#include "threadsafestringstack.h"

// -----------------------------------------------------------------------

int randInt(int min = 0, int max = 1000) {
    auto rg = QRandomGenerator::securelySeeded();
    return rg.bounded(min, max);
}

QString randStr(int len = 5) {
    const QString ltrs = "abcdefghijklmnopqrstuvwxyz0123456789";
    QString newStr; newStr.reserve(len);
    for (int i = 0; i < len; ++i) {
        const auto index =
                randInt(0, ltrs.size());
        newStr.push_back(ltrs[index]);
    }
    return newStr;
}

// -----------------------------------------------------------------------

void pusher(ThreadSafeStringStack& strs) {
    QThread::msleep(randInt());
    for (int i = 0; i < 10; ++i) {
        strs.push(randStr());
    }
}

void poper(ThreadSafeStringStack& strs) {
    QThread::msleep(randInt());
    for (int i = 0; i < 5; ++i) {
        strs.pop();
    }
}

// -----------------------------------------------------------------------

int main()
{
    const int thCount{ 100 };
    ThreadSafeStringStack strs;

    // ***

    QVector<QThread*> ths;
    for (int i = 0; i < thCount / 2; ++i) {
        ths.push_back(
                    QThread::create(
                        pusher, std::ref(strs)));
    }
    for (int i = 0; i < thCount / 2; ++i) {
        ths.push_back(
                    QThread::create(
                        poper, std::ref(strs)));
    }
    qDebug() << "Threads count:"
             << ths.size();

    // ***

    for (int i = 0; i < ths.size(); ++i) {
        ths[i]->start();
    }

    // ***

    for (int i = 0; i < ths.size(); ++i) {
        ths[i]->wait();

        delete ths[i];
        ths[i] = Q_NULLPTR;
    }

    qDebug() << strs.list();
    return 0;
}
