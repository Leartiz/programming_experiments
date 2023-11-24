#include <vector>
#include <iostream>

#include "spdlog/spdlog.h"
#include "spdlog/async.h"

#include "spdlog/sinks/stdout_sinks.h"
#include "spdlog/sinks/daily_file_sink.h"
#include "spdlog/sinks/rotating_file_sink.h"
#include "spdlog/sinks/basic_file_sink.h"

void reg_combined_logger()
{
    std::vector<spdlog::sink_ptr> sinks{
        std::make_shared<spdlog::sinks::stdout_sink_mt>(),
        std::make_shared<spdlog::sinks::daily_file_sink_mt>("logs/daily/log.txt", 23, 59),
        std::make_shared<spdlog::sinks::rotating_file_sink_mt>("logs/rotating/log.txt", 1024, 5),
        std::make_shared<spdlog::sinks::basic_file_sink_mt>("logs/basic/log.txt")
    };

    const auto combined_logger = std::make_shared<spdlog::async_logger>(
        "combined_logger", begin(sinks), end(sinks), spdlog::thread_pool()
        );

    // register it if you need to access it globally.
    spdlog::register_logger(combined_logger);
}

int main()
{
    spdlog::init_thread_pool(256, 1);
    reg_combined_logger();

    // ***

    auto logger = spdlog::get("combined_logger");
    if (!logger) {
        std::cout << "logger not exist\n";
        return 0;
    }

    for (int i = 0; i < 100; ++i) {
        logger->info("привет");
    }

    spdlog::shutdown();
    return 0;
}
