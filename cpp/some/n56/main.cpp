#include <vector>
#include <iostream>
#include <sstream>

struct Object final
{
    std::string val1;
    std::string val2;
    int number{ 0 };

    Object() = default;
    Object(const std::string& val1,
           const std::string& val2,
           int number)
        : val1{ val1 }
        , val2{ val2 }
        , number{ number } {}
};

void print_objects(std::vector<Object> objs)
{
    for (size_t i = 0; i < objs.size(); ++i) {
        std::cout << "val1: " << objs[i].val1;
        std::cout << "val2: " << objs[i].val2;
        std::cout << "number: " << objs[i].number;
        std::cout << std::endl;
    }
}

int main()
{
    std::stringstream text{
        "aedad dqwdwq 123\n"
        "dwqqq wqessw 341"
    };

    // *** push_back ***

    {
        text.seekg(std::ios_base::beg);
        std::vector<Object> objs;
        while (!text.eof()) {
            Object one;
            text >> one.val1 >> one.val2 >> one.number;
            objs.push_back(one);
        }

        print_objects(objs);
    }

    // *** emplace_back ***

    {
        text.seekg(std::ios_base::beg);
        std::vector<Object> objs;
        while (!text.eof()) {
            std::string val1;
            std::string val2;
            int number{ 0 };

            text >> val1 >> val2 >> number;

            objs.emplace_back(
                val1, val2, number);
        }

        print_objects(objs);
    }

    return 0;
}
