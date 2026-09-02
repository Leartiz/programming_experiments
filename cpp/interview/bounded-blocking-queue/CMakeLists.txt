cmake_minimum_required(VERSION 3.20)
project(bounded-blocking-queue LANGUAGES CXX)

add_executable(bounded-blocking-queue main.cpp)
target_compile_features(bounded-blocking-queue PRIVATE cxx_std_17)
