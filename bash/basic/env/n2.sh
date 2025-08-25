#!/bin/bash

print_func_name() {
    echo "FUNCNAME[0]: ${FUNCNAME[0]}"
    echo "FUNCNAME[1]: ${FUNCNAME[1]}"
    echo "FUNCNAME[2]: ${FUNCNAME[2]}"
    echo
}

# NOTE:
#   ${0} — имя самого скрипта или команды, которая запустила скрипт.
#   ${1} — первый аргумент, переданный скрипту.
#   ${2} — второй аргумент, переданный скрипту, и так далее.
#   ${#} — количество переданных аргументов.
#   ${@} или ${*} — все аргументы сразу.
#
print_positional_params() {
    echo "#: ${#}"
    echo "@: ${@}"
    echo "*: ${*}"

    echo "0: ${0}"
    echo "1: ${1}"
    echo "1: ${2}"
    echo
}

print_bash_source() {
    echo "BASH_SOURCE[0]: ${BASH_SOURCE[0]}"
    echo "BASH_SOURCE[1]: ${BASH_SOURCE[1]}"
    echo "BASH_SOURCE[2]: ${BASH_SOURCE[2]}"
    echo
}

print_bash_lineno() {
    echo "BASH_LINENO: ${BASH_LINENO}"
    echo
}

print_random() {
    echo "RANDOM $RANDOM"
    echo "RANDOM $RANDOM"
    echo "RANDOM $RANDOM"
    echo
}

print_command() {
    echo "BASH_COMMAND: ${BASH_COMMAND}"
    echo
}

# ------------------------------------------------------------------------

print_bash_source
print_positional_params
print_func_name
print_bash_lineno
print_random
print_command