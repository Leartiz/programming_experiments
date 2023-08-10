#include "myclass.h"

MyClass::MyClass(QObject* parent)
    : QObject{ parent }, m_readOnly{ false } {}

void MyClass::setReadOnly(bool ro) {
    m_readOnly = ro;
}

bool MyClass::isReadOnly() const {
    return m_readOnly;
}
