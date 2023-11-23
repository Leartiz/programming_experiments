#include <iostream>
#include <string>
#include <string_view>

struct Widget final
{
    explicit Widget(const std::string& value)
        : s{ value }  {
        std::cout << "Widget ctor\n";
    }
    ~Widget() {
        std::cout << "Widget dtor\n";
    }

    std::string s;
};

void consume(Widget w)
{
    std::cout << "consume";
    std::cout << "w.s: " << w.s << std::endl;
}

Widget doSomeVeryComplicatedThingWithSeveralArguments(
    int arg1, std::string_view arg2)
{
    static_cast<void>(arg1);
    static_cast<void>(arg2);
    //...

    return Widget{ "some" };
}

// cases
// -----------------------------------------------------------------------

void someFunc() {
    auto widget =
        doSomeVeryComplicatedThingWithSeveralArguments(
            123, "hello");
    consume(widget);
}

void someFunc1() {
    consume(
        doSomeVeryComplicatedThingWithSeveralArguments(
            123, "hello"));
}

int main()
{
    // copy-elision: rvo, nrvo.

    {
        std::cout << "someFunc: {\n";
        someFunc();
        std::cout << "}\n\n";
    }
    {
        std::cout << "someFunc1: {\n";
        someFunc1();
        std::cout << "}\n\n";
    }
    return 0;
}
