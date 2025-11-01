script_dir="$(
    cd -- "$(dirname -- "${BASH_SOURCE[0]}")" >/dev/null 2>&1
    pwd -P
)"
source "$script_dir/main.sh"

# NOTE:
#   перезаписать реализацию функции
: "[
colors::print_env() {
    echo "hi"
}
]"

setup_log_level_alias() {
    local orig_name="$1"
    local alias_name="$2"
    eval "${alias_name}() { ${orig_name} \"\$@\"; }"
}

# ------------------------------------------------------------------------

print_env() {
    sc::common::logging::print_env "$@";
}
logging::print_env() {
    sc::common::logging::print_env "$@";
}
common::logging::print_env() {
    sc::common::logging::print_env "$@";
}

# ------------------------------------------------------------------------

# sc::common::logging::print_env() {
#     echo "sc::common::logging::print_env (override)"
# }

#unset -f sc::common::logging::internal::print_env
