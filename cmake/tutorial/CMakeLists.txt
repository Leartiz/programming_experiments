cmake_minimum_required(VERSION 3.26)

project(tutorial)

message(STATUS "PROJECT_NAME: ${PROJECT_NAME}")

add_executable(${PROJECT_NAME}
    main.cpp
)
