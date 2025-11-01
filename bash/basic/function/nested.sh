sc::logging::log_level_str_to_num() {
    echo "call log_level_str_to_num"
}

sc::logging::log_level_num_to_str() {
    echo "call log_level_num_to_str"
}

sc::common::logging::internal::trim_namespace_prefix_and_alias() {
    echo "call trim_namespace_prefix_and_alias"
}

# to aliases
# ------------------------------------------------------------------------

sc::common::logging::using_namespace() {
    echo "call using_namespace"

    trim_alias() {
        sc::common::logging::internal::trim_namespace_prefix_and_alias "$@"
    }

    local remove_count="$1"
    echo "remove count: $remove_count"

    trim_alias sc::logging::log_level_str_to_num "$remove_count"
    trim_alias sc::logging::log_level_num_to_str "$remove_count"

    unset -f trim_alias
}

# ------------------------------------------------------------------------

sc::common::logging::using_namespace 2
trim_alias

#log_level_str_to_num
#log_level_num_to_str