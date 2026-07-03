#ifndef WITHPROPERTIES_H
#define WITHPROPERTIES_H

#include <QObject>

class WithProperties : public QObject
{
    Q_OBJECT

    Q_PROPERTY(int content MEMBER m_content)

public:
    WithProperties(QObject* parent = Q_NULLPTR);
    ~WithProperties() override;

    void printProps();

private:
    int m_content{ 0 };
};

#endif // WITHPROPERTIES_H
