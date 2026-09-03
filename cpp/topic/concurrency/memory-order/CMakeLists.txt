cmake_minimum_required(VERSION 3.20)
project(memory-order LANGUAGES CXX)

add_executable(memory-order main.cpp)
target_compile_features(memory-order PRIVATE cxx_std_17)

if(UNIX AND NOT APPLE)
  target_link_libraries(memory-order PRIVATE pthread)
endif()
