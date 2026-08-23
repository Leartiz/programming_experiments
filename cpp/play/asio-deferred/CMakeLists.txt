cmake_minimum_required(VERSION 3.16)
project(asio-deferred)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_package(Boost REQUIRED)
find_package(Threads REQUIRED)

add_executable(asio-deferred main.cpp)
target_include_directories(asio-deferred PRIVATE ${Boost_INCLUDE_DIRS})
target_link_libraries(asio-deferred PRIVATE Threads::Threads)
