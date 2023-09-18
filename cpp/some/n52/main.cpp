#include <QString>
#include <QDebug>

int main()
{
    {
        QString s1;
        QString s2{ "" };

        qDebug() << s1.isNull();
        qDebug() << s1.isEmpty();

        qDebug() << s2.isNull();
        qDebug() << s2.isEmpty();
    }

    {
        QString s{ "fff" };
        s.setNum(1001);

        qDebug() << s;
    }


}
