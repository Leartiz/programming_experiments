#include <iostream>

#include "../../dependencies/zpp/zpp_bits.h"

struct person
{
    std::string name;
    int age{};
};

struct person1
{
    int age{};
    std::string name;
};

int main()
{
    auto [data, in, out] = zpp::bits::data_in_out();

    // Serialize a few people:
    out(person{"Person1", 25}, person{"Person2", 35});


    // Define our people.
    person p1;
    person1 p2;

    // We can now deserialize them either one by one `in(p1)` `in(p2)`, or together, here
    // we chose to do it together in one line:
    in(p1, p2);


    std::cout << p2.age << " " << p2.name << "\n";
}
