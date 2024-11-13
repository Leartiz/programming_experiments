TEMPLATE = app
CONFIG += console c++17
CONFIG -= app_bundle
CONFIG -= qt

INCLUDEPATH += D:\education\my\pl_cpp\libs\boost_1_82_0
LIBS += -lws2_32 -lmswsock

SOURCES += main.cpp
