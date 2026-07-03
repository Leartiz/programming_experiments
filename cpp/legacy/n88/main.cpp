#include <iostream>
#include <stdexcept>

#include "spdlog/spdlog.h"

#include "spdlog/async.h"
#include "spdlog/sinks/rotating_file_sink.h"

void reg_async_logger()
{
    auto async_logger = spdlog::rotating_logger_mt<spdlog::async_factory>(
        "async", "logs/async_log.txt", 1024 * 1, 5);
    if (!async_logger)
        throw std::runtime_error("logger not created");


    async_logger->info("register");

    // ***

    if (async_logger->sinks().empty())
        throw std::runtime_error("logger has not sinks");

    auto rot_file_sink = std::static_pointer_cast<spdlog::sinks::rotating_file_sink_mt>(
        async_logger->sinks()[0]);

    if (rot_file_sink)
        std::cout << "logger file: " << rot_file_sink->filename() << '\n';
    std::cout << "logger name: " << async_logger->name() << '\n';
}

int main()
{
    try {
        reg_async_logger();
    }
    catch(std::runtime_error& err) {
        std::cerr << err.what() << '\n';
        return 0;
    }

    // ***

    auto logger = spdlog::get("async");
    if (!logger) {
        std::cout << "logger with name '" << "async" << "' not exist\n";
        return 0;
    }

    for (int i = 0; i < 10000; ++i) {
        logger->info(i);
    }

    spdlog::shutdown();
    return 0;
}
