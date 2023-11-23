#include <iostream>

#include <QDebug>
#include <QCoreApplication>
#include <QString>

using namespace std;

class Object{
public:
    Object(const std::string& value)
        : m_value{ value } { cout << "ctor " << m_value << "\n"; };
    ~Object() { cout << "dtor " << m_value << "\n"; }

    Object(const Object& obj) {
        if (&obj == this) return;
        cout << "copy " << m_value << "\n";
    }
    Object(Object&& obj) {
        if (&obj == this) return;
        cout << "move " << m_value << "\n";
    }

    Object& operator==(Object&& obj) {
        if (&obj == this) return *this;
        cout << "move assign" << m_value << "\n";
        return *this;
    }

    Object& operator==(const Object& obj) {
        if (&obj == this) return *this;
        cout << "copy assign" << m_value << "\n";
        return *this;
    }

    friend ostream& operator<<(ostream& out, const Object& obj) {
        return out << obj.m_value;
    };

    std::string m_value;
};


void from() {
    std::cout << "end\n\n\n";
}

template <typename... Args>
void from(const Object& number, Args&&... numbers) {
    std::cout << number << " ";
    from(numbers...);
}

int main()
{
    Object obj{ "555" };
    from(
        Object("123"), obj, Object("123"), obj, Object("123"), obj
        );
}
