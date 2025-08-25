script_dir="$(
    cd -- "$(dirname -- "${BASH_SOURCE[0]}")" >/dev/null 2>&1
    pwd -P
)"

#echo "script_dir: $script_dir"
source "$script_dir/n1_aliases.sh"

print_env XDG_CONFIG
colors::print_env XDG_CONFIG
