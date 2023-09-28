#include <string>
#include <iostream>
#include <stdexcept>

struct Exc
{
    explicit Exc(std::string& msg)
        : m_msg{ msg } {}
    virtual ~Exc() {}

public:
    virtual std::string msg() const {
        return m_msg;
    }

private:
    std::string m_msg;
};

struct ExcA_1 : public Exc
{

};

struct ExcA_2 : public Exc
{

};

struct ExcAA : ExcA_1, ExcA_2 {

};

// множественное наследование и исключения
// ромбовидное наследование и сключения

int main()
{

    return 0;
}
