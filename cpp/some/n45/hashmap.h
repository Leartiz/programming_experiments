#ifndef HASHMAP_H
#define HASHMAP_H

#include <stdexcept>

#include <QDebug>
#include <QTextStream>

#include <QVector>
#include <QString>
#include <QSharedPointer>
#include <QCryptographicHash>

namespace lez
{

template <typename Key>
QString hashFromKey(const Key& key)
{
    QString s;
    QTextStream ts{ &s };
    ts << key;

    // ***

    const auto hashBytes =
            QCryptographicHash::hash(
                s.toUtf8(), QCryptographicHash::Md5);

    return hashBytes.toHex();
}

size_t indexFromHash(const QString& hashValue)
{
    return qHash(hashValue);
}

// -----------------------------------------------------------------------

template <typename Key, typename T>
class HashMap
{
    struct Entry;
    using SpEntry = QSharedPointer<Entry>;

    struct Entry
    {
    public:
        Key key;
        T value;

    public:
        QString hashValue;
        SpEntry next;
    };

public:
    explicit HashMap(const size_t sz = 10);
    size_t sz() const;
    void clear();

public:
    void insert(const Key& k, const T& val);
    void set(const Key& k, const T& val);
    const T& get(const Key& k) const;
    void del(const Key& k);

public:
    QVector<QString> hashVals() const;
    QVector<Key> keys() const;
    QVector<T> vals() const;

public:
    QString toString() const;

private:
    QVector<SpEntry> m_vc;
    size_t m_size;
};

// -----------------------------------------------------------------------

template <typename Key, typename T>
HashMap<Key, T>::HashMap(const size_t sz)
    : m_size{ sz }
{
    m_vc.resize(m_size);
}

template <typename Key, typename T>
size_t HashMap<Key, T>::sz() const
{
    return m_size;
}

template <typename Key, typename T>
void HashMap<Key, T>::clear()
{
    m_vc.clear();
    m_vc.resize(m_size);
}

template <typename Key, typename T>
void HashMap<Key, T>::insert(const Key& k, const T& val)
{
    const QString hashValue = hashFromKey(k);
    const auto hashIndex = qHash(hashValue) % m_size;

    // ***

    const SpEntry sp = SpEntry(
                new Entry{
                    k, val, hashValue, nullptr
                });

    if (!m_vc[hashIndex]) {
        m_vc[hashIndex] = sp;
        return;
    }

    SpEntry curSp = m_vc[hashIndex];
    while (true) {
        if (curSp->key == k) {
            throw std::runtime_error{
                "duplicate key"
            };
        }

        if (curSp->next) {
            curSp = curSp->next;
        } else {
            break;
        }
    }
    curSp->next = sp;
}

template <typename Key, typename T>
void HashMap<Key, T>::set(const Key& k, const T& val)
{
    const QString hashValue = hashFromKey(k);
    const auto hashIndex = qHash(hashValue) % m_size;

    // ***

    const SpEntry sp = SpEntry(
                new Entry{
                    k, val, hashValue, nullptr
                });

    if (!m_vc[hashIndex]) {
        m_vc[hashIndex] = sp;
        return;
    }

    // ***

    SpEntry curSp = m_vc[hashIndex];
    while (true) {
        if (curSp->key == k) {
            curSp->value = val;
            return;
        }

        if (curSp->next) {
            curSp = curSp->next;
        } else {
            break;
        }
    }
    curSp->next = sp;
}

template <typename Key, typename T>
const T& HashMap<Key, T>::get(const Key& k) const
{
    const QString hashValue = hashFromKey(k);
    const auto hashIndex = qHash(hashValue) % m_size;

    // ***

    if (m_vc[hashIndex]) {
        auto cur = m_vc[hashIndex];
        while (cur) {
            if (cur.key == k) {
                return cur.value;
            }
            cur = cur->next;
        }
    }

    throw std::runtime_error{
        "no value by key"
    };
}

template <typename Key, typename T>
void HashMap<Key, T>::del(const Key& k)
{
    const QString hashValue = hashFromKey(k);
    const auto hashIndex = qHash(hashValue) % m_size;

    // ***

    if (m_vc[hashIndex]) {
        auto cur = m_vc[hashIndex];
        while (cur) {
            if (cur.key == k) {
                // TODO:

                return cur.value;
            }
            cur = cur->next;
        }
    }
}

// -----------------------------------------------------------------------

template <typename Key, typename T>
QVector<QString> HashMap<Key, T>::hashVals() const
{
    QVector<QString> result;
    for (int i = 0; i < m_vc.size(); ++i) {
        auto cur = m_vc[i];
        while (cur) {
            result.push_back(cur->hashValue);
            cur = cur->next;
        }
    }
    return result;
}

template <typename Key, typename T>
QVector<Key> HashMap<Key, T>::keys() const
{
    QVector<Key> result;
    for (int i = 0; i < m_vc.size(); ++i) {
        auto cur = m_vc[i];
        while (cur) {
            result.push_back(cur->key);
            cur = cur->next;
        }
    }
    return result;
}

template <typename Key, typename T>
QVector<T> HashMap<Key, T>::vals() const
{
    QVector<T> result;
    for (int i = 0; i < m_vc.size(); ++i) {
        auto cur = m_vc[i];
        while (cur) {
            result.push_back(cur->value);
            cur = cur->next;
        }
    }
    return result;
}

template <typename Key, typename T>
QString HashMap<Key, T>::toString() const
{
    QString res;
    QTextStream ts{ &res };
    for (int i = 0; i < m_vc.size(); ++i) {
        ts << "# " << i << ": ";
        auto cur = m_vc[i];
        while (cur) {
            ts << cur->value << " ";
            cur = cur->next;
        }
        Qt::endl(ts);
    }
    return res;
}

}

#endif // HASHMAP_H
