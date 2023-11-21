#include <QDebug>
#include <QSharedPointer>
#include <QJsonDocument>

#include "userdto.h"

int main()
{
    UserDto* user = new UserDto();
    user->Username = "12123";
    qDebug() << user->toJson();

    // ***

    QJsonDocument jsonDoc = QJsonDocument::fromJson(
        R"({"FirstName":"1name","SecondName":"123","Username":"12123", "Ages":[123,123,123,123]})"
    );

    UserDto* user1 = new UserDto();
    qDebug() << jsonDoc.object();

    user1->fromJson(jsonDoc.object());

    qDebug() << "Username:" << user1->Username;
    qDebug() << "Username:" << user1->FirstName;
    qDebug() << "Username:" << user1->SecondName;

    qDebug() << "Ages:" << user1->Ages;

    return 0;
}
