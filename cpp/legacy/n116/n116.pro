TEMPLATE = app
CONFIG += console c++17
CONFIG -= app_bundle
CONFIG -= qt

BOOST_PATH = $$getenv(BOOST_PATH)
message("BOOST_PATH = $$BOOST_PATH")

INCLUDEPATH += $$BOOST_PATH
LIBS += -lws2_32 -lmswsock

SOURCES += main.cpp
