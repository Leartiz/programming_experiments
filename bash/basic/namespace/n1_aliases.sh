script_dir="$(
    cd -- "$(dirname -- "${BASH_SOURCE[0]}")" >/dev/null 2>&1
    pwd -P
)"

echo "script_dir: $script_dir"
source "$script_dir/n1.sh"

# NOTE:
#   перезаписать реализацию функции
: `[
colors::print_env() {
    echo "hi"
}
]`

print_env() {
    colors::print_env "$@";
}

print_env