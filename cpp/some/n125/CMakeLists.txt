cmake_minimum_required(VERSION 3.5)

project(n125 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

set(THIRD_PARTY_DIR "${CMAKE_SOURCE_DIR}/third_party")
add_subdirectory(${THIRD_PARTY_DIR}/usrsctp)
include_directories(${THIRD_PARTY_DIR}/usrsctp/usrsctplib)

set(SRCS
    src/main.cpp
)

add_executable(n125 ${SRCS})
target_link_libraries(n125 usrsctp)

include(GNUInstallDirs)
install(TARGETS n125
    LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
    RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR}
)
