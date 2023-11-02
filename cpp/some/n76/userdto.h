#ifndef USERDTO_H
#define USERDTO_H

#include "dto.h"
#include "roomdto.h"

class UserDto : public DataTransferObject
{
    Q_GADGET
    QT_REDECL_STATIC_META_OBJECT()

    QT_DECL_PROPERTY(QString, Username)
    QT_DECL_PROPERTY(QString, FirstName)
    QT_DECL_PROPERTY(QString, SecondName)
    QT_DECL_PROPERTY_WITH_VALUE(int, Age, 100)
    QT_DECL_PROPERTY(RoomDto, Room)
};

#endif // USERDTO_H
