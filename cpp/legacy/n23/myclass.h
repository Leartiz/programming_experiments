#ifndef MYCLASS_H
#define MYCLASS_H

#include <QObject>

class MyClass : public QObject
{
    Q_OBJECT

public:
    explicit MyClass(QObject* parent = nullptr);

public:
    Q_INVOKABLE void initWithException();
    Q_INVOKABLE void init();

signals:
    void initFailed(QString msg);
};

#endif // MYCLASS_H
