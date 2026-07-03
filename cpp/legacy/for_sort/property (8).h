#ifndef PROPERTY_H
#define PROPERTY_H

#include <QMap>
#include <QVariant>
#include <typeinfo>

class BaseProperty {
public:
    BaseProperty() {}
    virtual ~BaseProperty() = default;

    virtual QVariant* data() = 0;
    virtual const std::type_info& getTypeInfo() = 0;

    auto getValue() {
        return *data();
    }

    virtual void setValue(const QVariant& val) = 0;
};

class PropertyRepository {
protected:
    QMap<QString, BaseProperty*> properties;

public:
    void insertProperty(const QString& key, BaseProperty* value) {
        properties.insert(key, value);
    }

    BaseProperty* property(const QString& key) {
        return properties[key];
    }

    template<class T>
    void setProperty(const QString& key, T value)
    {
        auto property = properties[key];

        property->setValue(value);
    }
};

class Property: public BaseProperty
{
    QString m_key;
    int m_typeId;
    QVariant value;

public:
    Property(QString name, QVariant val, PropertyRepository* repository):
        m_key(name), value(val)
    {
        m_typeId = value.typeId();

        repository->insertProperty(m_key, this);
    }

    const QString key() {
        return m_key;
    }

    QVariant* data() override {
        return static_cast<QVariant*>(&value);
    }

    void setValue(const QVariant& val) override {

        qDebug() << m_typeId << val.typeId();

        Q_ASSERT(m_typeId == val.typeId());

        value.setValue(val);
    }

    const std::type_info& getTypeInfo() override {
        return typeid(value);
    }

};

#define DECL_PROPERTY(type, name, value) Property name = Property(QString(#name), static_cast<type>(value), this);


class Dto: public PropertyRepository
{
    // Property<qint64> id = Property<qint64>("id", 1, &properties);
public:

    DECL_PROPERTY(qint64, id, 1)
    DECL_PROPERTY(QString, name, "OK")

    qint64 dosomething() {
        id.setValue(qint64(3));

        return id.getValue().toLongLong();
    }
};

#endif // PROPERTY_H
