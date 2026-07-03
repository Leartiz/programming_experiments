QT = core

CONFIG += c++17 cmdline

# You can make your code fail to compile if it uses deprecated APIs.
# In order to do so, uncomment the following line.
#DEFINES += QT_DISABLE_DEPRECATED_BEFORE=0x060000    # disables all the APIs deprecated before Qt 6.0.0

SOURCES += \
        ../../dependencies/quill/src/LogLevel.cpp \
        ../../dependencies/quill/src/PatternFormatter.cpp \
        ../../dependencies/quill/src/Quill.cpp \
        ../../dependencies/quill/src/Utility.cpp \
        ../../dependencies/quill/src/bundled/fmt/format.cc \
        ../../dependencies/quill/src/bundled/fmt/os.cc \
        ../../dependencies/quill/src/detail/HandlerCollection.cpp \
        ../../dependencies/quill/src/detail/LoggerCollection.cpp \
        ../../dependencies/quill/src/detail/SignalHandler.cpp \
        ../../dependencies/quill/src/detail/ThreadContext.cpp \
        ../../dependencies/quill/src/detail/ThreadContextCollection.cpp \
        ../../dependencies/quill/src/detail/backend/BackendWorker.cpp \
        ../../dependencies/quill/src/detail/backend/BacktraceStorage.cpp \
        ../../dependencies/quill/src/detail/backend/StringFromTime.cpp \
        ../../dependencies/quill/src/detail/backend/TimestampFormatter.cpp \
        ../../dependencies/quill/src/detail/backend/TransitEventBuffer.cpp \
        ../../dependencies/quill/src/detail/misc/FileUtilities.cpp \
        ../../dependencies/quill/src/detail/misc/Os.cpp \
        ../../dependencies/quill/src/detail/misc/RdtscClock.cpp \
        ../../dependencies/quill/src/detail/misc/Utilities.cpp \
        ../../dependencies/quill/src/handlers/ConsoleHandler.cpp \
        ../../dependencies/quill/src/handlers/FileHandler.cpp \
        ../../dependencies/quill/src/handlers/Handler.cpp \
        ../../dependencies/quill/src/handlers/JsonFileHandler.cpp \
        ../../dependencies/quill/src/handlers/RotatingFileHandler.cpp \
        ../../dependencies/quill/src/handlers/StreamHandler.cpp \
        main.cpp

# Default rules for deployment.
qnx: target.path = /tmp/$${TARGET}/bin
else: unix:!android: target.path = /opt/$${TARGET}/bin
!isEmpty(target.path): INSTALLS += target

DISTFILES += \
    ../../dependencies/quill/src/bundled/fmt/LICENSE.rst

HEADERS += \
    ../../dependencies/quill/include/quill/Clock.h \
    ../../dependencies/quill/include/quill/Config.h \
    ../../dependencies/quill/include/quill/Fmt.h \
    ../../dependencies/quill/include/quill/LogLevel.h \
    ../../dependencies/quill/include/quill/Logger.h \
    ../../dependencies/quill/include/quill/MacroMetadata.h \
    ../../dependencies/quill/include/quill/PatternFormatter.h \
    ../../dependencies/quill/include/quill/Quill.h \
    ../../dependencies/quill/include/quill/QuillError.h \
    ../../dependencies/quill/include/quill/TransitEvent.h \
    ../../dependencies/quill/include/quill/TweakMe.h \
    ../../dependencies/quill/include/quill/Utility.h \
    ../../dependencies/quill/include/quill/bundled/fmt/args.h \
    ../../dependencies/quill/include/quill/bundled/fmt/chrono.h \
    ../../dependencies/quill/include/quill/bundled/fmt/color.h \
    ../../dependencies/quill/include/quill/bundled/fmt/compile.h \
    ../../dependencies/quill/include/quill/bundled/fmt/core.h \
    ../../dependencies/quill/include/quill/bundled/fmt/format-inl.h \
    ../../dependencies/quill/include/quill/bundled/fmt/format.h \
    ../../dependencies/quill/include/quill/bundled/fmt/os.h \
    ../../dependencies/quill/include/quill/bundled/fmt/ostream.h \
    ../../dependencies/quill/include/quill/bundled/fmt/printf.h \
    ../../dependencies/quill/include/quill/bundled/fmt/ranges.h \
    ../../dependencies/quill/include/quill/bundled/fmt/std.h \
    ../../dependencies/quill/include/quill/bundled/fmt/xchar.h \
    ../../dependencies/quill/include/quill/clock/TimestampClock.h \
    ../../dependencies/quill/include/quill/detail/HandlerCollection.h \
    ../../dependencies/quill/include/quill/detail/LogMacros.h \
    ../../dependencies/quill/include/quill/detail/LogManager.h \
    ../../dependencies/quill/include/quill/detail/LoggerCollection.h \
    ../../dependencies/quill/include/quill/detail/LoggerDetails.h \
    ../../dependencies/quill/include/quill/detail/Serialize.h \
    ../../dependencies/quill/include/quill/detail/SignalHandler.h \
    ../../dependencies/quill/include/quill/detail/ThreadContext.h \
    ../../dependencies/quill/include/quill/detail/ThreadContextCollection.h \
    ../../dependencies/quill/include/quill/detail/backend/BackendWorker.h \
    ../../dependencies/quill/include/quill/detail/backend/BacktraceStorage.h \
    ../../dependencies/quill/include/quill/detail/backend/StringFromTime.h \
    ../../dependencies/quill/include/quill/detail/backend/TimestampFormatter.h \
    ../../dependencies/quill/include/quill/detail/backend/TransitEventBuffer.h \
    ../../dependencies/quill/include/quill/detail/misc/Attributes.h \
    ../../dependencies/quill/include/quill/detail/misc/Common.h \
    ../../dependencies/quill/include/quill/detail/misc/FileUtilities.h \
    ../../dependencies/quill/include/quill/detail/misc/Os.h \
    ../../dependencies/quill/include/quill/detail/misc/Rdtsc.h \
    ../../dependencies/quill/include/quill/detail/misc/RdtscClock.h \
    ../../dependencies/quill/include/quill/detail/misc/TypeTraitsCopyable.h \
    ../../dependencies/quill/include/quill/detail/misc/Utilities.h \
    ../../dependencies/quill/include/quill/detail/spsc_queue/BoundedQueue.h \
    ../../dependencies/quill/include/quill/detail/spsc_queue/UnboundedQueue.h \
    ../../dependencies/quill/include/quill/filters/FilterBase.h \
    ../../dependencies/quill/include/quill/handlers/ConsoleHandler.h \
    ../../dependencies/quill/include/quill/handlers/FileHandler.h \
    ../../dependencies/quill/include/quill/handlers/Handler.h \
    ../../dependencies/quill/include/quill/handlers/JsonFileHandler.h \
    ../../dependencies/quill/include/quill/handlers/NullHandler.h \
    ../../dependencies/quill/include/quill/handlers/RotatingFileHandler.h \
    ../../dependencies/quill/include/quill/handlers/StreamHandler.h
