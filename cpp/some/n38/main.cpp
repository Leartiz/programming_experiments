#include <future>
#include <thread>
#include <chrono>
#include <iostream>

using namespace std;

std::future<bool> submit_form(const std::string& form)
{
    const auto handle =
            [](const std::string& form) -> bool {

        std::this_thread::sleep_for(
                    std::chrono::milliseconds(
                        500));

        return
            form == "famous";
    };

    // ***

    std::packaged_task<bool(const std::string&)> task(handle);
    auto future = task.get_future();

    std::thread thread(std::move(task), form);
    thread.detach();

    return future;
}

int main()
{
    auto f0 = submit_form("famous");
    auto f1 = submit_form("unknown");

    // ***

    std::cout << f0.get() << std::endl;
    std::cout << f1.get() << std::endl;

    return 0;
}
