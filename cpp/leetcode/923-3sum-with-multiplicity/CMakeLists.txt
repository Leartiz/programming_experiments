cmake_minimum_required(VERSION 3.16)
project(three_sum_with_multiplicity LANGUAGES CXX)

add_executable(three_sum_with_multiplicity main.cpp)
add_executable(combinations_test combinations_test.cpp)

set_property(TARGET three_sum_with_multiplicity PROPERTY CXX_STANDARD 17)
set_property(TARGET combinations_test PROPERTY CXX_STANDARD 17)
