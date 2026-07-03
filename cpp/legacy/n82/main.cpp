#include <iostream>

class A final {
public:
    A() {}
    ~A() try {
        throw std::runtime_error{"something broke"};
    }
    catch (const std::runtime_error& err) {
        std::cerr << "A dtor err:" << err.what() << std::endl;

        return;

        // implicit ---
        //            |
        //   -----------
        //   |
        //   V
        // throw;
    }
};

int main()
{
    try {
        const auto a = A();
        static_cast<void>(a);
    }
    catch(const std::runtime_error& err) {
        std::cerr << "main err: " << err.what() << std::endl;
    }

    std::cout << "[OK]" << std::endl;
}
