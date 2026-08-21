cmake_minimum_required(VERSION 3.16)
project(flat-map)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_package(Boost REQUIRED)

add_executable(flat-map main.cpp)
target_include_directories(flat-map PRIVATE ${Boost_INCLUDE_DIRS})
