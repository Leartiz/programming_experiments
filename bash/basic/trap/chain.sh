script_dir="$(
    cd -- "$(dirname -- "${BASH_SOURCE[0]}")" >/dev/null 2>&1
    pwd -P
)"
source "$script_dir/_common.sh"

# ------------------------------------------------------------------------

add_trap() {
    local new_cmd="$1"
    local signal="$2"

    echo $signal

    local current_cmd=$(trap -p "$signal" | sed -E "s/trap -- '(.*)' $signal/\1/")
    echo "current_cmd: ${current_cmd}"

    if [[ -z "$current_cmd" ]]; then
        trap "$new_cmd" "$signal"
    else
        trap "$current_cmd; $new_cmd" "$signal"
    fi
}

add_trap 'echo "Первый обработчик"' ERR

current_cmd=$(trap -p "ERR" | sed -E "s/trap -- '(.*)' $signal/\1/")
echo "current_cmd: ${current_cmd}"

add_trap 'echo "Второй обработчик"' ERR

count=0
while true ; do
    sleep 0.25s
    count=$((count + 1))
    echo "count: $count"

    echo "generated: $(generate_random)"

    generate_random_and_check 7
done