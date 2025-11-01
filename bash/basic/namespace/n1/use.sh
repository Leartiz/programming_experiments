script_dir="$(
    cd -- "$(dirname -- "${BASH_SOURCE[0]}")" >/dev/null 2>&1
    pwd -P
)"
source "$script_dir/main_aliases.sh"

print_env XDG_CONFIG
logging::print_env XDG_CONFIG
common::logging::print_env XDG_CONFIG
sc::common::logging::print_env XDG_CONFIG

sc::common::logging::internal::print_env XDG_CONFIG