cmake_minimum_required(VERSION 3.10)

project(n134 CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

message("PROJECT NAME: " + ${PROJECT_NAME})

add_executable(${PROJECT_NAME}
    main.cpp
)