#ifndef ROOMDTO_H
#define ROOMDTO_H

#include "dto.h"

class RoomDto : public DataTransferObject
{
    Q_GADGET
public:
    QT_DECL_PROPERTY(QString, Name)
    QT_DECL_PROPERTY(double, Temperature)

protected:
    QMetaObject metaObject() const override {
        return staticMetaObject;
    }
};

#endif // ROOMDTO_H
