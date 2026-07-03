#include <iostream>
#include <vector>
#include <thread>
#include <chrono>

using namespace std;

class Widget final
{
public:
    void foo()
    {
        std::cout << "Widget.foo()" << std::endl;
    }

    void const_foo() const
    {
        std::cout << "Widget.foo()" << std::endl;
    }
};

struct Point final
{
public:
    int x{ 0 }, y{ 0 };
};

bool f(const Widget& w)
{
    decltype(w) ww = w;
    ww.const_foo();
    //ww.foo();

    return true;
}

// -----------------------------------------------------------------------

// trailing return type!
//
template <typename Container, typename Index>
auto authAndAccess(Container& c, Index i) -> decltype(c[i])
{
    // some actions...

    return c[i];
}

template <typename Container, typename Index>
auto authAndAccessWithAuto(Container& c, Index i)
{
    return c[i];
}

template <typename Container, typename Index>
decltype(auto) authAndAccessWithDecltypeAuto(Container& c, Index i)
{
    return c[i];
}

// -----------------------------------------------------------------------

int main()
{
    {
        Widget w;
        const Widget& cw = w;

        auto myWidget1 = cw; // Widget
        myWidget1.foo();

        // decltype type inference rules!
        //
        decltype(auto) myWidget2 = cw; // const Widget&
        //myWidget2.foo();
        myWidget2.const_foo();
    }
    {
        std::vector<int> vec;
        vec.resize(100);

        vec[0] = 1;
        std::cout << "vec[0]: " << vec[0] << std::endl;
        std::cout << "authAndAccess(vec, 0): " << authAndAccess(vec, 0) << std::endl;

        authAndAccess(vec, 0) = 100;
        std::cout << "vec[0]: " << vec[0] << std::endl;

        //authAndAccessWithAuto(vec, 0) = 200;

        authAndAccessWithDecltypeAuto(vec, 0) = 200;
        std::cout << "vec[0]: " << vec[0] << std::endl;
    }
    {
        const int i = 0;
        decltype(i) ii = 0; // const int
        //ii = 1;
    }
    {
        Widget w;
        f(w);

        const decltype(f)& ff = f;
        ff(w);

        const auto auto_ff = &f;
        auto_ff(w);
    }
    {
        const decltype(Point::x) x = 100;
        const decltype(Point::y) y = 100;

        std::cout << "x = " << x << std::endl;
        std::cout << "y = " << y << std::endl;
    }
    {
        std::vector<int> v;
        decltype(v)& vv = v;
        std::cout << "&v = " << &v << std::endl;
        std::cout << "&vv = " << &vv << std::endl;

        vv.resize(1);
        vv[0] = 100;

        std::cout << "vv[0] = " << vv[0] << std::endl;
        std::cout << "v[0] = " << v[0] << std::endl;

        int i = 0;
        decltype(v[0]) ref_i = i;
        std::cout << "i = " << i << std::endl;
        std::cout << "ref_i = " << ref_i << std::endl;

        ref_i = 100;

        std::cout << "i = " << i << std::endl;
        std::cout << "ref_i = " << ref_i << std::endl;
    }

    {
        using namespace chrono;
        std::this_thread::sleep_for(100ms); // ?
    }
    return 0;
}
