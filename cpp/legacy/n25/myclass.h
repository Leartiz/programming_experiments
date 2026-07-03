#ifndef MYCLASS_H
#define MYCLASS_H

#include <QObject>

class MyClass : public QObject
{
    Q_OBJECT

public:
    explicit MyClass(QObject *parent = nullptr);

signals:
    void innerValuesChanged();

public slots:
    void onInnerValuesChanged();
};

#endif // MYCLASS_H
