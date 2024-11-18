#include <mutex>
#include <iostream>

using namespace std;

int main()
{
    {
        std::mutex mx;
        std::unique_lock<std::mutex> ll{ mx };
        //...
    }
    {
        std::recursive_mutex rmx;
        std::unique_lock<
            std::recursive_mutex> ll{ rmx };

        //ll.lock(); // don't work!

        rmx.lock();
        rmx.lock();
        rmx.lock();

        //...
    }

    std::cout << "[OK]" << '\n';
    return 0;
}
