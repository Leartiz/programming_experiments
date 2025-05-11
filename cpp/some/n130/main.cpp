#include <string>
#include <iostream>
#include <utility>
#include <thread>

using namespace std;

struct Message final
{
    std::string first;
    std::string second;

    int condition;
};

void try_send(Message&& msg);

void send(Message&& msg)
{
    std::cout << "[send] msg addr: " << &msg << std::endl;
    for (int i = 0; i < 3; ++i) {
        msg.condition = i;
        try_send(std::move(msg));

        std::cout << "msg.first: " << msg.first << std::endl;
        std::cout << "msg.second: " << msg.second << std::endl;
    }
}

void try_send(Message&& msg)
{
    std::cout << "[try_send] msg addr: " << &msg << std::endl;
    if (msg.condition == 2) {
        auto other_msg = std::move(msg);
        static_cast<void>(other_msg);
    }
}

int main()
{
    int* data = new int[100];

    std::cout << "[main beg]" << std::endl;
    Message msg {
        .first = "first",
        .second = "second",
        .condition = 0
    };

    std::cout << "[main before send] msg addr: " << &msg << std::endl;
    send(std::move(msg));
    std::cout << "[main after send] msg addr: " << &msg << std::endl;

    std::this_thread::sleep_for(std::chrono::milliseconds(1000));
    std::cout << "[main end]" << std::endl;

    delete[] data;
}
