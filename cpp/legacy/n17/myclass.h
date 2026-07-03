#ifndef MYCLASS_H
#define MYCLASS_H

#include <QObject>

class MyClass : public QObject
{
    Q_OBJECT

    Q_PROPERTY(bool readOnly
               READ isReadOnly
               WRITE setReadOnly
               NOTIFY readOnlyChanged)

public:
    explicit MyClass(QObject* parent);

public:
    void setReadOnly(bool ro);
    bool isReadOnly() const;

signals:
    void readOnlyChanged(bool val);

private:
    bool m_readOnly{ false };
};
#endif // MYCLASS_H
