function myfunc {
    echo "This is an example of using a function"
}

myfunc

greet() {
    local name="${1:-Гость}"     # \$1 — первый аргумент; значение по умолчанию
    local lang="${2:-ru}"        # \$2 — второй аргумент
    echo "name=$name lang=$lang"
}

greet "Алиса" "en"   # name=Алиса lang=en
greet

log() {
    local lvl="\$1"; shift
    local n=${LVL[$lvl]:-3}

    if [ "$CUR" -ge "$n" ]; then
        local ts
        ts=$(date +"%H:%M:%S")

        local color=""
        if [ "${LOG_COLOR:-1}" = "1" ]; then
            case "$lvl" in
                debug) color="\e[2m";;
                info)  color="\e[36m";;
                warn)  color="\e[33m";;
                error) color="\e[31m";;
            esac
        fi

        echo -e "${color}[$ts][$CI_JOB_NAME][$lvl] $* \e[0m"
    fi
}

log info "LOG_LEVEL=${LOG_LEVEL}"