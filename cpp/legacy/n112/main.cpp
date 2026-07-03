#include <iostream>
#include <memory>
#include <array>
#include <iterator>

struct ZF_HEAD {
    int a;
    int b;

    ~ZF_HEAD() {
        std::cout << "~ZF_HEAD()" << std::endl;
    }
};

std::array<int, 256> aaa{5, 6, 7};
std::shared_ptr<ZF_HEAD> GetSpZfHead() {
    return std::shared_ptr<ZF_HEAD>(
        reinterpret_cast<ZF_HEAD*>(aaa.data()));
}

ZF_HEAD* GetRawPtrZfHead() {
    return reinterpret_cast<ZF_HEAD*>(aaa.data());
}

int main()
{
    {
        auto zf = GetRawPtrZfHead();
        std::cout << "zf->a: " << zf->a << std::endl;
        std::cout << "zf->b: " << zf->b << std::endl;
        zf->b = 1000;

        std::copy(aaa.begin(), aaa.end(),
                  std::ostream_iterator<int>(std::cout, ","));

    }

//    {
//        auto zf = GetSpZfHead();
//        std::cout << "zf->a: " << zf->a << std::endl;
//        std::cout << "zf->b: " << zf->b << std::endl;
//    }

    std::cout << "[OK]" << std::endl;
    return 0;
}
