#include <string>
#include <iostream>
#include <stdexcept>

struct Exc
{
    explicit Exc(const std::string& msg)
        : m_msg{ msg } {}
    virtual ~Exc() {}

public:
    virtual std::string msg() const {
        return m_msg;
    }

private:
    std::string m_msg;
};

// -----------------------------------------------------------------------

struct ExcA_1 : /*virtual*/ Exc
{
    using Exc::Exc;
};

struct ExcA_2 : /*virtual*/ Exc
{
    using Exc::Exc;
};

struct ExcAA : ExcA_1, ExcA_2
{
    explicit ExcAA(const std::string& value)
        :
          /*Exc{ value + "_base" },*/

          ExcA_1{ value + "_1" } // ?
        ,
          ExcA_2{ value + "_2" } // ?
    {}
};

// -----------------------------------------------------------------------

void fuu_throws_exc() {
    throw Exc{ "fuu_throws_exc" };
}

void fuu_throws_exc_aa() {
    throw ExcAA{ "fuu_throws_exc" };
}

int main()
{
    try {
        fuu_throws_exc();
    }
    catch (const Exc& e) {
        std::cout << e.msg() << "\n";
    }

    // ***

    try {
        fuu_throws_exc_aa();
    }
    catch (const ExcAA& e) {
        std::cout << e.ExcA_1::msg() << "\n";
        std::cout << e.ExcA_2::msg() << "\n";
    }

    // ***

    try {
        fuu_throws_exc_aa();
    }
    catch (const ExcA_1& e) {
        std::cout << e.msg() << "\n";
    }

    // ***

    try {
        fuu_throws_exc_aa();
    }
    catch (const ExcA_2& e) {
        std::cout << e.msg() << "\n";
    }

    // ***
    std::cout << "\n";

    try {
        fuu_throws_exc_aa();
    }
    catch (const Exc& e) {
        std::cout << "e.msg: " << e.msg() << "\n";

        auto excAA = dynamic_cast<const ExcAA&>(e);
        std::cout << "excAA.ExcA_1::msg: " << excAA.ExcA_1::msg() << "\n";
        std::cout << "excAA.ExcA_2::msg: " << excAA.ExcA_2::msg() << "\n";
        //std::cout << "excAA.msg: " << excAA.msg() << "\n";
    }

    return 0;
}
