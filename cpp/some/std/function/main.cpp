#include <iostream>
#include <functional>

struct PrintNum
{
    void operator()(int i) const
    {
        std::cout << i << '\n';
    }
};

int main() {
    std::function<void(int)> f_display_obj = PrintNum();
    f_display_obj(18);
}