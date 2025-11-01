sc::common::logging::internal::declare_global_vars() {
    SC_COMMON_LOGGING_COLLAPSE_SECTIONS=1
    SC_COMMON_LOGGING_ENABLE_COLOR=1
    SC_COMMON_LOGGING_ENABLE_TRAP=1
    SC_COMMON_LOGGING_CONSOLE_LOG_LEVEL="debug"
}

print_global_vars() {
    echo "COLLAPSE_SECTIONS: ${SC_COMMON_LOGGING_COLLAPSE_SECTIONS}"
    echo "ENABLE_COLOR:      ${SC_COMMON_LOGGING_ENABLE_COLOR}"
    echo "ENABLE_TRAP:       ${SC_COMMON_LOGGING_ENABLE_TRAP}"
    echo "CONSOLE_LOG_LEVEL: ${SC_COMMON_LOGGING_CONSOLE_LOG_LEVEL}"
}

declare_tmp_var() {
    local use_local=${1:-1}

    local tmp="tmp value"
    echo $tmp
}

declare_tmp_var
echo "tmp=${tmp}"

# ------------------------------------------------------------------------

sc::common::logging::internal::declare_global_vars
print_global_vars

echo

