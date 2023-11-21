#ifndef DTO_H
#define DTO_H

#include <QObject>
#include <QJsonObject>
#include <QSharedPointer>

// -----------------------------------------------------------------------

#define QT_DECL_REQUIRED_PROPERTY(type, key) \
Q_PROPERTY(type key MEMBER key) \
    type key;

#define QT_DECL_REQUIRED_PROPERTY_WITH_VALUE(type, key, defaultValue) \
Q_PROPERTY(type key MEMBER key) \
    type key{ defaultValue };

#define QT_DECL_OPTIONAL_PROPERTY(type, key) \
Q_PROPERTY(QVariant key MEMBER m_##key) \
private: \
    QVariant m_##key; \
public: \
    type key() { return m_##key.value<type>(); } \
    void reset##key() { m_##key = QVariant::fromValue(nullptr); } \
    void set##key(type value) { m_##key = QVariant::fromValue(value); }

// -----------------------------------------------------------------------

#define QT_DECL_REQUIRED_DTO_AS_PROPERTY(type, key) \
Q_PROPERTY(QSharedPointer<DataTransferObject> key MEMBER m_##key) \
private: \
    QSharedPointer<DataTransferObject> m_##key{ new type{} }; \
public: \
    QSharedPointer<type> key() {  return qSharedPointerCast<type>(m_##key); };

#define QT_DECL_OPTIONAL_DTO_AS_PROPERTY(type, key) \
Q_PROPERTY(QSharedPointer<DataTransferObject> key MEMBER m_##key) \
private: \
    QSharedPointer<DataTransferObject> m_##key{ new type{} }; \
public: \
    QSharedPointer<type> key() { return qSharedPointerCast<type>(m_##key); } \
    void set##key(type* value) { m_##key.reset(value); }

#define QT_DECL_REQUIRED_PROPERTY_LIST(type, key) \
    Q_PROPERTY(QVariantList key WRITE set##key READ get##key) \
        QList<type> key; \
    \
    void set##key(QVariantList list) { \
        key.clear(); \
        for (const auto& el : list) { \
            key.push_back(qvariant_cast<type>(el)); \
        } \
    } \
    \
    QVariantList get##key() { \
        QVariantList out; \
        for (const auto& el : key) { \
            out.push_back(QVariant::fromValue(el)); \
        } \
        return out; \
    }

// -----------------------------------------------------------------------

#define QT_DECL_REQUIRED_DTO_LIST_AS_PROPERTY(type, key) \
    Q_PROPERTY(QList<QSharedPointer<DataTransferObject>> key MEMBER m_##key) \
        private: \
    QList<QSharedPointer<DataTransferObject>> m_##key; \
        public: \
        QList<QSharedPointer<type>> key() { \
            QList<QSharedPointer<type>> result; \
                for (int i = 0; i < m_##key.size(); ++i) { \
                result.push_back( \
                    qSharedPointerCast<type>(m_##key[i])); \
            } \
            return result; \
        }; \
    QList<QSharedPointer<DataTransferObject>>& ref##key() { \
        return m_##key; \
    }

// -----------------------------------------------------------------------

#define QT_REDECL_STATIC_META_OBJECT() \
protected: \
    QMetaObject metaObject() const override { \
        return staticMetaObject; \
    }

#define QT_DECL_DTO() \
private: \
    Q_GADGET \
    QT_REDECL_STATIC_META_OBJECT() \
public:

// -----------------------------------------------------------------------

class DataTransferObject {
    struct DTOParsingError : public std::runtime_error {
        using std::runtime_error::runtime_error;
    };

private:
    using DTOPointer = QSharedPointer<DataTransferObject>;
    using DTOPointerList = QList<QSharedPointer<DataTransferObject>>;

private:
    Q_GADGET

public:
    virtual QJsonObject toJson() const;
    virtual void fromJson(const QJsonObject&);
    virtual ~DataTransferObject() = default;

protected:
    void fromJsonList(const QMetaProperty& property, const QJsonArray& array);
    void fromJsonObject(const QMetaProperty& property, const QJsonObject& object);
    virtual QMetaObject metaObject() const = 0;
};

#endif // DTO_H
