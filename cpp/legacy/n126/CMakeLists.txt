cmake_minimum_required(VERSION 3.10)

project(n126)
message(STATUS "PROJECT_NAME: ${PROJECT_NAME}")
message(STATUS "CMAKE_INSTALL_PREFIX: ${CMAKE_INSTALL_PREFIX}")

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

### Find out build type ###
set(CMAKE_CONFIGURATION_TYPES Debug Release)
if(NOT CMAKE_BUILD_TYPE)
    set(CMAKE_BUILD_TYPE Release)
endif()
message("Current configuration type: ${CMAKE_BUILD_TYPE}")

set(CMAKE_CXX_FLAGS "-g3 -Wall -Wextra -Wunreachable-code -pedantic")
set(CMAKE_CXX_FLAGS_DEBUG "-O0 -fno-inline -pg")
set(CMAKE_CXX_FLAGS_RELEASE "-O3 -DNDEBUG")

### Boost ###
set(Boost_COMPONENTS system log log_setup)
set(Boost_USE_STATIC_LIBS ON)
set(Boost_USE_MULTITHREADED ON)
set(Boost_USE_STATIC_RUNTIME OFF)
find_package(Boost 1.75 COMPONENTS ${Boost_COMPONENTS} REQUIRED)

### Threads ###
set(CMAKE_THREAD_PREFER_PTHREAD TRUE)
find_package(Threads REQUIRED)

### Ppconsul ###
find_package(ppconsul REQUIRED)

set(SOURCE_EXE main.cpp)
add_executable(${PROJECT_NAME} ${SOURCE_EXE})

target_include_directories(${PROJECT_NAME} PRIVATE
    ${Boost_INCLUDE_DIR}
)
message(STATUS "PPCONSUL_INCLUDE_DIRS: ${PPCONSUL_INCLUDE_DIRS}")
message(STATUS "Boost_INCLUDE_DIR: ${Boost_INCLUDE_DIR}")
message(STATUS "Boost_LIBRARIES: ${Boost_LIBRARIES}")

target_link_libraries(${PROJECT_NAME}
    ${Boost_LIBRARIES}
    ppconsul
)