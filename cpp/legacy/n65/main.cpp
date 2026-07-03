#include <string>
#include <memory>
#include <iostream>
#include <functional>

struct A final
{
    std::string str;
    int value;

    ~A() {
        std::cout << "~A()\n";
    }
};

struct S final
{
    A* value{
        nullptr
    };

    ~S() {
        std::cout << "~S()\n";
    }
};

int main()
{
    {
        S s;
        s.value = new A();
        //delete s.value;
    }

    // ***

    std::unique_ptr<S, std::function<void(S*)>> ps =
            std::make_unique<S>();

    ps->value = new A{ "123", 123 };
    ps.get_deleter() = [](S* thiz) {
        delete thiz->value;
        thiz->value = NULL;
        delete thiz;
    };

    return 0;
}
