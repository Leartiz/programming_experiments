#ifndef ROOMDTO_H
#define ROOMDTO_H

#include "dto.h"

class RoomDto : public DataTransferObject
{
    Q_GADGET
    QT_REDECL_STATIC_META_OBJECT()

    QT_DECL_PROPERTY(QString, Name)
    QT_DECL_PROPERTY(double, Temperature)
};

#endif // ROOMDTO_H
