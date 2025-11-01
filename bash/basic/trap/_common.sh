#!/bin/bash

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
