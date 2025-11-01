#!/bin/bash

foo() {
    trap 'echo "Trap установлен в foo"' ERR
    local count=0
    echo "count: $count"
    ((count++))
}

bar() {
    trap 'echo "Trap установлен в bar"' ERR
    local count=0
    echo "count: $count"
}

dee() {
    local count=0
    echo "count: $count"
    ((count++))
}

# NOTE:
#   не срабатывает
#     а точнее просто не перезаписывает trap, который был в функции foo
#
trap 'echo "Ошибка в строке ${LINENO}: команда ${BASH_COMMAND}"; exit 1' ERR

# NOTE:
#   странное поведение вызовов
#     trap-обработчиков
#
foo
bar

bar
foo

count=0
echo "count: $count"
((count++))

count=101
echo "count: $count"

dee