#include <iostream>
#include <stdexcept>

class A {
public:
    explicit A(int a) : m_a{ a } {
        std::cout << "A\n";
    }
    ~A() {
        std::cout << "~A\n";
        try {
            cleaner(m_a);
        }
        catch(...) {
            std::cerr << "A dtor ...\n";
        }
    }
private:
    void cleaner(int a) {
        if (a < 0) {
            throw a;
        }
    }

    int m_a;
};

void fn(int a)
{
    A objA{ a };
    static_cast<void>(objA);
    //...
    throw std::runtime_error{
        "something broke"
    };
}

int main()
{
    try {
        fn(-100);
    }
    catch(int a) {
        std::cerr << "main " << a << "\n";
    }
    catch(const std::runtime_error& err) {
        std::cerr << "main " << err.what() << "\n";
    }
    catch(...) {
        std::cerr << "main ...\n";
    }

    std::cout << "[OK]\n";
    return 0;
}
