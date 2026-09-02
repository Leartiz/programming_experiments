cmake_minimum_required(VERSION 3.20)
project(thread-safe-cache LANGUAGES CXX)

add_executable(thread-safe-cache main.cpp)
target_compile_features(thread-safe-cache PRIVATE cxx_std_17)
