#!/bin/bash

# NOTE:
#   https://tokmakov.msk.ru/blog/item/96

generate_random() {
    echo $(( RANDOM % 100 ))
}

generate_random_and_check() {
    local divisor=$1
    local rand_num=$(generate_random)

    if (( rand_num % divisor == 0 )); then
        echo "Число $rand_num делится на $divisor без остатка — вызываем ошибку"
        return 1 # ERR
    else
        echo "Число $rand_num не делится на $divisor без остатка — все нормально"
        return 0 # OK
    fi
}

# ------------------------------------------------------------------------

# NOTE:
#   ERR срабатывает, когда в скрипте происходит ошибка,
#     команда возвращает ненулевой код завершения.
#

echo "[$0] vs. [${BASH_SOURCE[0]}]"

trap 'echo "Ошибка в строке ${LINENO}: команда ${BASH_COMMAND}"; exit 1' ERR

trap 'echo "Получен сигнал SIGINT (Ctrl+C), завершаем скрипт"; exit 1' SIGINT

count=0
while true ; do
    sleep 0.25s
    count=$((count + 1))
    echo "count: $count"

    echo "generated: $(generate_random)"

    generate_random_and_check 7
done