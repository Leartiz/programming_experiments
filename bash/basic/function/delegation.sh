#!/bin/dash

# define
# ------------------------------------------------------------------------

# NOTE:
#   valid function name characters: [a-zA-Z_][a-zA-Z0-9_]*
#
sc__common__sections__section_start() {
    section_name=${1}
    section_desc=${2}
    open=${3}

    echo "section name=${section_name}"
    echo "section desc=${section_desc}"
    echo "open=${open}"
    echo
}

sc__common__sections__collapsed_section_start() {
    sc__common__sections__section_start "$@" "false"
}

# use
# ------------------------------------------------------------------------

sc__common__sections__section_start \
    "name" "very long desc" "true"

sc__common__sections__collapsed_section_start  \
    "name" "very long desc" "true"

# ------------------------------------------------------------------------

sc__common__sections__collapsed_section_start  \
    "name" "very long desc"

sc__common__sections__collapsed_section_start  \
    "name" "very long desc" "true"

# ------------------------------------------------------------------------

sc__common__sections__collapsed_section_start  \
    "name" "true"

sc__common__sections__collapsed_section_start  \
    "name" "" "true"
