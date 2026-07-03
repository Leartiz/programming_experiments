#include <csignal>
#include <iostream>
#include <thread>

int main()
{
    std::signal(SIGINT, [](int code){
        std::cout << code << " SIGINT" << '\n';
        std::this_thread::sleep_for(std::chrono::seconds(3));
    });
    std::signal(SIGABRT, [](int code){
        std::cout << code << " SIGABRT" << '\n';
        std::this_thread::sleep_for(std::chrono::seconds(3));
    });
    std::signal(SIGTERM, [](int code){
        std::cout << code << " SIGTERM" << '\n';
        std::this_thread::sleep_for(std::chrono::seconds(3));
    });
    std::signal(SIGBREAK, [](int code){
        std::cout << code << " SIGBREAK" << '\n';
        std::this_thread::sleep_for(std::chrono::seconds(3));
    });
    std::signal(SIGILL, [](int code){
        std::cout << code << " SIGILL" << '\n';
        std::this_thread::sleep_for(std::chrono::seconds(3));
    });

    while(true);

    std::cout << "ret";
    return 0;
}
