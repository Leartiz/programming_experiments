#include <QJsonArray>

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
        auto variant = metaproperty.readOnGadget(this);

        if (variant.canConvert<DTOPointer>()) {
            auto dto = variant.value<DTOPointer>();
            if (dto) variant = dto->toJson();
            else
                continue;
        }
        else if (variant.canConvert<DTOPointerList>()) {
            auto dtos = variant.value<DTOPointerList>();
            if (!dtos.isEmpty()) {
                QJsonArray values;
                for (int i = 0; i < dtos.size(); ++i)
                    values.push_back(dtos[i]->toJson());
                variant = values;
            }
            else
                continue;
        }

        result.insert(
            key,
            QJsonValue::fromVariant(variant)
            );
    }
    return result;
}
void DataTransferObject::fromJson(const QJsonObject& object)
{
    QMetaObject mo = metaObject();

    qDebug() << object;

    for(auto it = object.constBegin(); it < object.constEnd(); it++)
    {
        const auto key = it.key();
        const auto value = it.value();

        auto index = mo.indexOfProperty(key.toStdString().c_str());


        if(index == -1)
            throw DTOParsingError("Property with key " + key.toStdString() + " not found.");

        auto prop = mo.property(index);

        if(value.isArray()) // List
        {
            fromJsonList(prop, value.toArray());
        }

        if(value.isObject()) // DTO
        {
            fromJsonObject(prop, value.toObject());
        }

        prop.writeOnGadget(static_cast<void*>(this), value.toVariant());
    }
}

void DataTransferObject::fromJsonList(const QMetaProperty& property, const QJsonArray& array)
{
    QVariantList list;

    for(auto it = array.cbegin(); it < array.cend(); it++)
    {
        if(it->isArray())
            throw DTOParsingError("Unsupported feature: arrays in arrays.");

        if(it->isObject())
            throw DTOParsingError("Unsupported feature: objects in arrays.");

        list.append(it->toVariant());
    }

    qDebug() << "list:" << list;

    property.writeOnGadget(static_cast<void*>(this), list);
}

void DataTransferObject::fromJsonObject(const QMetaProperty& property, const QJsonObject& object)
{
    auto prop_meta = property.enclosingMetaObject();

    for(const auto& key: object.keys())
    {
        if(object.value(key).isArray())
            throw DTOParsingError("Unsupported feature: arrays in objects.");

        if(object.value(key).isObject())
            throw DTOParsingError("Unsupported feature: objects in objects.");

        const auto index = prop_meta->indexOfProperty(key.toStdString().c_str());

        if(index == -1)
            throw DTOParsingError("Property with key " + key.toStdString() + " not found.");

        auto current = prop_meta->property(index);
        current.writeOnGadget(&current, object.value(key));
    }
}
