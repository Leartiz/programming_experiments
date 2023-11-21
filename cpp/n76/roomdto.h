#ifndef ROOMDTO_H
#define ROOMDTO_H

#include "dto.h"

class RoomDto : public DataTransferObject
{
    QT_DECL_DTO()

public:
    QT_DECL_REQUIRED_PROPERTY(QString, Name)
    QT_DECL_REQUIRED_PROPERTY(double, Temperature)
};

#endif // ROOMDTO_H
