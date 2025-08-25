colors::print_env() {
    local prefix="$1"
    env | grep "^${prefix}" | sort
}
