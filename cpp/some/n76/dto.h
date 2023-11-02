#ifndef DTO_H
#define DTO_H

#include <QObject>
#include <QJsonObject>

// -----------------------------------------------------------------------

#define QT_DECL_PROPERTY(type, key) \
Q_PROPERTY(type key MEMBER key) \
    type key;

#define QT_DECL_DTO_AS_PROPERTY(type, key) \
Q_PROPERTY(type key MEMBER key) \
    type key{ new type{} };

#define QT_DECL_PROPERTY_WITH_VALUE(type, key, defaultValue) \
    Q_PROPERTY(type key MEMBER key) \
        type key{ defaultValue };

#define QT_REDECL_STATIC_META_OBJECT() \
protected: \
    QMetaObject metaObject() const override { \
        return staticMetaObject; \
    } \
public:

// -----------------------------------------------------------------------

class DataTransferObject {
    Q_GADGET
public:
    virtual QJsonObject toJson() const;
    virtual void fromJson(const QJsonObject&);
    virtual ~DataTransferObject() = default;

protected:
    virtual QMetaObject metaObject() const {
        return staticMetaObject;
    };
};

#endif // DTO_H
