#!/bin/bash

script_dir="$(
    cd -- "$(dirname -- "${BASH_SOURCE[0]}")" >/dev/null 2>&1
    pwd -P
)"
source "$script_dir/_common.sh"

# ------------------------------------------------------------------------

foo() {
    local count=0
    echo "count (before): $count"
    ((count++))
    ((count++))
    ((count++))
    echo "count (after): $count"

    return 1 # ERR
}

LOGGING_ENABLE_TRAP=1
echo "LOGGING_ENABLE_TRAP: ${LOGGING_ENABLE_TRAP}"

if [[ $LOGGING_ENABLE_TRAP -eq 1 ]]; then
    echo "Set Trap"
    trap 'echo "Ошибка в строке ${LINENO}: команда ${BASH_COMMAND}"; exit 1' ERR
else
    echo "No Trap"
fi

foo