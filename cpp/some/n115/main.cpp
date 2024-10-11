#include <iostream>
#include <memory>

struct Point {
    int x, y, z;
};

struct Vertex {
    Point points[3];
};

void createVerties(size_t p)
{
    static_cast<void>(p);

    std::shared_ptr<Vertex> ss(new Vertex);
    //...
}

int main()
{
    std::cout << "sizeof(Point): " << sizeof(Point) << '\n';
    std::cout << "sizeof(Vertex): " << sizeof(Vertex) << '\n';

    return 0;
}
