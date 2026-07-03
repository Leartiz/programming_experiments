#include <QtCore>
#include <QtConcurrent/QtConcurrent>

QString tooUpper(const QString& str)
{
    qDebug() << "Current thread id:"
             << QThread::currentThreadId();

    return str.toUpper();
}

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);

    {
        QFuture<QString> future =
                QtConcurrent::run(
                    tooUpper, QString("test"));

        future.waitForFinished();
        qDebug() << future.result();
    }

    // ***

    {
        QStringList strs;
        strs << "one" << "two" << "three";
        QFuture<QString> future =
                QtConcurrent::mapped(
                    strs.begin(), strs.end(),
                    tooUpper);

        future.waitForFinished();
        qDebug() << future.results();
    }

    return a.exec();
}
