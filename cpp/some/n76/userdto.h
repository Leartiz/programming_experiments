#ifndef USERDTO_H
#define USERDTO_H

#include "dto.h"
#include "roomdto.h"

class UserDto : public DataTransferObject
{
    Q_GADGET
public:
    QT_DECL_PROPERTY(QString, Username)
    QT_DECL_PROPERTY(QString, FirstName)
    QT_DECL_PROPERTY(QString, SecondName)
    QT_DECL_PROPERTY(int, Age)
    QT_DECL_PROPERTY(RoomDto*, Room)

protected:
    QMetaObject metaObject() const override {
        return staticMetaObject;
    }
};

#endif // USERDTO_H
