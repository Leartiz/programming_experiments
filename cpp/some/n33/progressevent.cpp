#include "progressevent.h"

ProgressEvent::ProgressEvent()
    : QEvent{ static_cast<Type>(Progress) } {}

void ProgressEvent::setValue(int val)
{
    m_value = val;
}

int ProgressEvent::value() const
{
    return m_value;
}
