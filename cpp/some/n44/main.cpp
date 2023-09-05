#include <QDebug>

#include <QHash>
#include <QMultiHash>
#include <QString>

struct MyClass
{
    QString firstName;
    QString secondName;
};

bool operator==(const MyClass& mc1, const MyClass& mc2)
{
    return
            mc1.firstName == mc2.firstName &&
            mc1.secondName == mc2.secondName;
}

uint qHash(const MyClass& mc)
{
    return
            qHash(mc.firstName) ^
            qHash(mc.secondName);
}

int main()
{
//    {
//        QHash<QString, MyClass> hash;
//    }

    {
        QHash<MyClass, QString> hash;
        hash[MyClass{ "first", "second" }]   = "412453";
        hash[MyClass{ "first1", "second1" }] = "756341";

        qDebug() << hash[MyClass{ "first", "second" }];
        qDebug() << hash[MyClass{ "first1", "second1" }];
    }

    return 0;
}
