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

        if (value.canConvert<DataTransferObject*>()) {
            const auto dto = value.value<DataTransferObject*>();
            if (dto) value = QVariant::fromValue(dto->toJson());
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
