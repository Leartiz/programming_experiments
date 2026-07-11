cmake_minimum_required(VERSION 3.16)
project(partition_array_into_three_parts LANGUAGES CXX)

add_executable(partition_array_into_three_parts main.cpp)

set_property(TARGET partition_array_into_three_parts PROPERTY CXX_STANDARD 17)
