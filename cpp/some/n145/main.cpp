#include <iostream>

template <typename Derived>
struct base
{
    void bf() { 
        static_cast<Derived*>(this)->df(); 
    }
};

template <typename T>
struct derived;

template <typename T = derived<bool>>
struct derived : public base<T>
{
    void df() { std::cout << "-- df\n"; }
};

int main()
{
    derived().bf();
    return 0;
}