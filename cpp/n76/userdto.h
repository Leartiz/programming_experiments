#ifndef USERDTO_H
#define USERDTO_H

#include "dto.h"
#include "roomdto.h"

class UserDto : public DataTransferObject
{
    QT_DECL_DTO()
    QT_DECL_REQUIRED_PROPERTY(QString, Username)
    QT_DECL_REQUIRED_PROPERTY(QString, FirstName)
    QT_DECL_REQUIRED_PROPERTY(QString, SecondName)
    QT_DECL_REQUIRED_PROPERTY_LIST(int, Ages)
    QT_DECL_REQUIRED_PROPERTY(QStringList, Addresses)
};

#endif // USERDTO_H
