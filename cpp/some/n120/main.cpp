#include <iostream>
#include <thread>

void run_thread() {
    std::thread t([]{
        while (true) {
            std::this_thread::sleep_for(std::chrono::seconds(1));
            std::cout << "." << '\n';
        }
    });
    t.detach();
}

int main()
{
    run_thread();
    std::chrono::seconds(5);
    return 0;
}
