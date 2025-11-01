#!/bin/dash

say_hello() {
    echo "hello $1"
}

say_bye() {
    echo "bye $1"
}

fn="hello"
full="say_${fn}"

# ------------------------------------------------------------------------

# NOTE:
#   Display information about command type.
#
if type "$full" >/dev/null 2>&1; then
    "$full" "world"
else
    printf 'No such function: %s\n' "$full" >&2
fi

# NOTE:
#   Execute a simple command or display information about commands.
#   -v    print a description of COMMAND similar to the `type' builtin
#
if command -v "$full" >/dev/null; then
    "$full" "world"
else
    printf 'No such function or command: %s\n' "$full" >&2
fi

# NOTE:
#   bash
#
# if declare -F "$full" >/dev/null; then
#   "$full" "world"
# else
#   printf 'No such function: %s\n' "$full" >&2
# fi