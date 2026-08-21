#include <boost/asio.hpp>
#include <chrono>
#include <iostream>

namespace asio = boost::asio;

int main() {
    asio::io_context ioc;
    asio::steady_timer timer(ioc, std::chrono::milliseconds(50));

    auto op = timer.async_wait(asio::deferred);

    std::move(op)([](boost::system::error_code ec) {
        std::cout << (ec ? ec.message() : "timer done") << '\n';
    });

    ioc.run();
}
