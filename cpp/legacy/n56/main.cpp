#include <vector>
#include <iostream>
#include <sstream>
#include <string>

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

    const std::string to_str() const {
        std::stringstream sout;
        sout << "val1: "   << val1 << ", ";
        sout << "val2: "   << val2 << ", ";
        sout << "number: " << number;
        return sout.str();
    }

    static Object create(std::istream& in) {
        Object one;
        in
                >> one.val1
                >> one.val2
                >> one.number;

        return one;
    }
};

void print_objects(std::vector<Object> objs)
{
    for (size_t i = 0; i < objs.size(); ++i) {
        std::cout << objs[i].to_str();
        std::cout << std::endl;
    }
}

int main()
{
    std::istringstream text{
        "aedad dqwdwq 123\n"
        "dwqqq wqessw 341"
    };

    // *** push_back ***

    {
        text.seekg(std::ios_base::beg);
        std::vector<Object> objs;
        while (!text.eof()) {
            objs.push_back(
                        Object::create(
                            text));
        }

        print_objects(objs);
    }

    // *** emplace_back ***

    {
        text.seekg(std::ios_base::beg);
        std::vector<Object> objs;
        while (!text.eof()) {
            std::string val1; //
            std::string val2; // --- object insides
            int number{ 0 };  //

            text
                    >> val1
                    >> val2
                    >> number;

            objs.emplace_back(
                val1, val2, number);
        }

        print_objects(objs);
    }

    return 0;
}
