setup_log_alias_by_prefix_count() {
    local full_name="$1"
    local remove_count="$2"

    local oldIFS=$IFS
    IFS=':'

    # NOTE:
    #   -r          do not allow backslashes to escape any characters
    #
    #   -a array    assign the words read to sequential indices of the array
    #               variable ARRAY, starting at zero
    #
    read -ra parts <<< "${full_name//::/:}"

    local parts_count=${#parts[@]}
    local start_index=$remove_count
    if (( start_index >= parts_count )); then
        start_index=$((parts_count - 1))
    fi

    local new_name="${parts[start_index]}"
    for (( i=start_index+1; i<parts_count; i++ )); do
        new_name+="::${parts[i]}"
    done

    eval "${new_name}() { ${full_name} \"\$@\"; }"
    IFS=$oldIFS
}

# visual test
# ------------------------------------------------------------------------

sc::common::logging::log_level_str_to_num() {
    echo "Original function called with $@"
}

setup_log_alias_by_prefix_count sc::common::logging::log_level_str_to_num 1
setup_log_alias_by_prefix_count sc::common::logging::log_level_str_to_num 2
setup_log_alias_by_prefix_count sc::common::logging::log_level_str_to_num 3
setup_log_alias_by_prefix_count sc::common::logging::log_level_str_to_num 100

sc::common::logging::log_level_str_to_num "test0"
common::logging::log_level_str_to_num "test1"
logging::log_level_str_to_num "test2"
log_level_str_to_num "test3"