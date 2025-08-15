# NOTE:
#   https://tokmakov.msk.ru/blog/item/90

greet() {
    echo $#

    shift 2
    echo $#
    local name="${1:-NotProvided}"

    shift
    echo $#
    local lang="${1:-NotProvided}"
    echo "name=$name lang=$lang"
}

greet Skipped Skipped Known En
greet Skipped Known En
greet Known En