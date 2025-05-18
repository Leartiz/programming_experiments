#include <string>
#include <iostream>
#include <sstream>
#include <utility>

// ----------------------------------------------

struct Widget final
{
    std::string str_value;
    int int_value{ 0 };
};

std::string to_string(const Widget& w)
{
    std::ostringstream sout;
    sout << "widget: { " << "\n";
    sout << " str_value: " << w.str_value << ",\n";
    sout << " int_value: " << w.str_value << "\n";
    sout << "}";
    return sout.str();
}

// ----------------------------------------------

void foo_lvalue(Widget& w) 
{
    std::cout << "addr(cw): " << &w << std::endl;
    std::cout << to_string(w) << std::endl;
}

void foo_rvalue(Widget&& w) 
{
    std::cout << "addr(w): " << &w << std::endl;
    foo_lvalue(w);

    Widget _ = std::move(w);
    (void)_;

    foo_lvalue(w);
}

// ----------------------------------------------

int main() 
{
    Widget w{ "111", 111 };
    
    std::cout << " *** " << std::endl;
    foo_rvalue(std::move(w));
    std::cout << std::endl;

    std::cout << " *** " << std::endl;
    foo_rvalue(Widget{ "555", 555 });
    std::cout << std::endl;
    
    return 0;
}