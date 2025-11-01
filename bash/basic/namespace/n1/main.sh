sc::common::logging::print_env() {
    sc::common::logging::internal::print_env
}

sc::common::logging::internal::print_env() {
    echo "sc::common::logging::internal::print_env()"

    local prefix="$1"
    env | grep "^${prefix}" | sort
}

