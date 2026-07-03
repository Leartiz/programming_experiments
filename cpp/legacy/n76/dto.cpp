#include <QMetaObject>
#include <QMetaProperty>

#include "dto.h"

QJsonObject DataTransferObject::toJson() const
{
    QJsonObject result;
    QMetaObject mo = metaObject();
    const auto count = mo.propertyCount();
    for (int i = mo.propertyOffset(); i < count; ++i) {
        const auto metaproperty = mo.property(i);

        const auto key = metaproperty.name();
        auto value = metaproperty.readOnGadget(this);

        if (value.canView(mo.metaType())) {
            void* unknownDto = value.data();
            auto dto = reinterpret_cast<DataTransferObject*>(unknownDto);
            value = QVariant::fromValue(dto->toJson());
        }


        result.insert(
            key,
            QJsonValue::fromVariant(value)
            );
    }
    return result;
}

void DataTransferObject::fromJson(const QJsonObject& value)
{

}
