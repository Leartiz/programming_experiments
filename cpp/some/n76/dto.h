#ifndef DTO_H
#define DTO_H

#include <QObject>
#include <QJsonObject>

#define QT_DECL_PROPERTY(type, key) \
    Q_PROPERTY(type key MEMBER key) \
        type key;

// -----------------------------------------------------------------------

class DataTransferObject {
    Q_GADGET
public:
    virtual QJsonObject toJson() const;
    virtual void fromJson(const QJsonObject&);
    virtual ~DataTransferObject() = default;

protected:
    virtual QMetaObject metaObject() const = 0;
};

#endif // DTO_H
