cmake_minimum_required(VERSION 3.10)

project(n139 CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_package(GTest REQUIRED)

message("PROJECT NAME: ${PROJECT_NAME}")

enable_testing()

add_executable(
    ${PROJECT_NAME}
    main.cpp
)

target_link_libraries(
    ${PROJECT_NAME}
    GTest::GTest GTest::gmock GTest::Main
)

include(GoogleTest)
gtest_discover_tests(${PROJECT_NAME})