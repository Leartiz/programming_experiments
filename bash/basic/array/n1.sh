
print_arr() {
    local input_arr=${1}

    echo "elements: " ${arr[*]}
    echo "length: " ${#arr[*]}
    echo "indexes: " ${!arr[*]}
}

# -----------------------------------------------

test_arr_0() {
    local arr=(Hello World)
    arr[2]="!"

    print_arr arr
}

test_arr_1() {
    local arr=(one two three four five)
    echo "arr[4]: " ${arr[4]}

    print_arr arr
}

test_arr_2() {
    local arr=(trace debug info warn warning error err fatal)

    print_arr arr
}

test_arr_0
test_arr_1
test_arr_2

# -----------------------------------------------

is_val_in_arr_0() {
    local input_level="${1}"

    local result="false"
    local log_levels=(trace debug info warn warning error err fatal)

    for cur_level in "${log_levels[@]}"; do
        if [ "$cur_level" = "$input_level" ]; then
            result="true"
            break
        fi
    done
    echo ${result}
}

is_val_in_arr_1() {
    local input_level="${1}"

    local result="false"
    local log_levels=(trace debug info warn warning error err fatal)

    for cur_level in "${log_levels[@]}"; do
        if [ "$cur_level" = "$input_level" ]; then
            return 0 # OK!
        fi
    done
    return 1
}

echo "$(is_val_in_arr_0 "info")"

if is_val_in_arr_1 "info"; then
    echo "true"
else
    echo "false"
fi

if ! is_val_in_arr_1 "info"; then
    echo "false"
else
    echo "true"
fi
