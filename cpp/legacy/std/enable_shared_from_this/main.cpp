#include <iostream>
#include <memory>

class Foo : public std::enable_shared_from_this<Foo>
{
public:
    virtual std::shared_ptr<Foo> getptr()
    {
        return shared_from_this();
    }
};

class Bar : public Foo 
{
public:
    std::shared_ptr<Foo> getptr() override
    {
        return shared_from_this();
    }
};
 
int main() {


}