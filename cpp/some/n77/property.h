#ifndef PROPERTY_H
#define PROPERTY_H

#include <QString>
#include <QVariant>

#include "baseproperty.h"

//class PropertyRepository;
//class Property : public BaseProperty
//{
//    QString m_key;

//public:
//    QVariant value;

//    Property(QString name, QVariant val, PropertyRepository* repository):
//        m_key(name), value(val)
//    {
//        repository->insertProperty(m_key, this);
//    }

//    const QString key() {
//        return m_key;
//    }

//    QVariant* data() override {
//        return static_cast<QVariant*>(&value);
//    }

//    const std::type_info& getTypeInfo() override {
//        return typeid(value);
//    }
//};

#endif // PROPERTY_H
