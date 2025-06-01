#include <functional>
#include <iostream>
#include <memory>

class Widget final
{
public:
    explicit Widget(int value = 0)
        : value{ value }
    {}

public:
    static std::unique_ptr<Widget> create_ptr(const int value = 0)
    {
        return std::make_unique<Widget>(value);
    }

public:
    int value{ 0 };
};

bool less(const std::unique_ptr<Widget>& lhs,
          const std::unique_ptr<Widget>& rhs)
{
    return lhs->value < rhs->value;
}

int main()
{
    {
        int x1;
        auto x2 = 0;
        //auto x3;
    }
    {
        // C++14
        auto lambda_less = [](const auto& lhs, const auto& rhs) {
            return lhs < rhs;
        };

        std::cout << lambda_less(1, 2) << std::endl;
        std::cout << lambda_less(2.0, 1.2) << std::endl;
        std::cout << lambda_less(std::string{ "123" },
                                 std::string{ "223" }) << std::endl;
    }
    {
        std::function<bool(const std::unique_ptr<Widget>&,
                           const std::unique_ptr<Widget>&)> local_less = less;

        const auto w1 = Widget::create_ptr(5);
        const auto w2 = Widget::create_ptr(10);

        std::cout << "local_less(w1, w2): " << local_less(w1, w2) << std::endl;
        std::cout << "local_less(w2, w1): " << local_less(w2, w1) << std::endl;
    }
    {
        std::function<void(int)> f;
        if (f) {
            f(100); // <--- !!!
        }
    }
}