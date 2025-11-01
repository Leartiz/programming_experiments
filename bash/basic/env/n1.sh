#!/bin/dash

# ------------------------------------------------------------------------

print_env_variables_exclude_prefix() {
    local prefix="$1"
    env | grep -v "^${prefix}" | sort
}

print_env_variables_include_prefix() {
    local prefix="$1"
    env | grep "^${prefix}" | sort
}

# ------------------------------------------------------------------------

print_env_variables_by_prefixes() {
    include_mode=1
    exclude_mode=2

    mode=$include_mode
    if [ "$1" = "--include" ]; then
        mode=$include_mode
        shift
    elif [ "$1" = "--exclude" ]; then
        mode=$exclude_mode
        shift
    fi

    if [ "$#" -eq 0 ]; then
        prefixes=""
    else
        prefixes="$*"
    fi

    for prefix in $prefixes; do
        if [ -n "$grep_expr" ]; then
            grep_expr="$grep_expr|^$prefix"
        else
            grep_expr="^$prefix"
        fi
    done

    if [ "$mode" -eq 1 ]; then
        env | grep -E "$grep_expr" | sort
    else
        env | grep -v -E "$grep_expr" | sort
    fi
}

# ------------------------------------------------------------------------

print_env_variables_by_prefixes --include "XDG_SESSION" "TERM"
echo

# print_env_variables_by_prefixes --exclude "XDG_SESSION" "TERM"
# echo

# print_env_variables_by_prefixes
# echo

# ------------------------------------------------------------------------