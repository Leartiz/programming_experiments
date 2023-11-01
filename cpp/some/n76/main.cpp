#include <QDebug>

#include "userdto.h"

int main()
{
    UserDto* user = new UserDto();
    user->Username = "12123";
    user->Room = new RoomDto();
    user->Room->Name = "123";

    qDebug() <<  user->toJson();
    qDebug() << user->Age;


    return 0;
}
