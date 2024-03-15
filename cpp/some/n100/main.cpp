#include <iostream>
#include <string>

class Foo
{
public:
    template<typename T>
    const Foo& fn(const T& value) const {
        std::cout << "typeid: " << typeid(T).name() << " ";
        std::cout << "value: " << value << "\n";
        return *this;
    }
};

int main()
{
    Foo f;
    f
        .fn(100)
        .fn(std::string("string"))
        .fn("chars");

    return 0;
}
