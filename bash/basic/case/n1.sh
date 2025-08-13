# test
# ------------------------------------------------------------------------

switch_num() {
    local num="${1:-1}"
    case $num in
        1)
            echo "num = 1"
        ;;
        2)
            echo "num = 2"
        ;;
        3)
            echo "num = 3"
        ;;
    esac
}

# sections
# ------------------------------------------------------------------------

function section_start() {
    local section_title=${1}
    local section_description=${2:-$section_title}

    local section_collapsed=${3:-1}
    local section_collapsed_str=""

    case "${section_collapsed}" in
        1|true|yes|on) section_collapsed_str="true" ;;
        *)             section_collapsed_str="false" ;;
    esac

    echo "section_start:`date +%s`:${section_title}[collapsed=${section_collapsed_str}]\r\e[0K${section_description}"
}

function section_end() {
  local section_title="${1}"

  echo -e "section_end:`date +%s`:${section_title}\r\e[0K"
}

# log
# ------------------------------------------------------------------------

# NOTE:
#   уровни взяты из `sc_logger`
#
# LOG_ANY   0
# LOG_FATAL 1
# LOG_ERROR 2
# LOG_WARN  3
# LOG_INFO  4
# LOG_DEBUG 5
# LOG_TRACE 6
#

__level_to_num() {
  case "${1:-info}" in
    trace) echo 6 ;;
    debug) echo 5 ;;

    info)         echo 4 ;;
    warning|warn) echo 3 ;;

    error|err) echo 2 ;;
    fatal)     echo 1 ;;
    *) echo 3 ;; # может быть не указано
  esac
}

lvl_strings=(trace debug info warn err fatal)
for lvl_str in "${lvl_strings[@]}"; do
    echo "lvl_num: $lvl_str = $(__level_to_num $lvl_str)"
done

log_line() {
    local lvl=${1:-"info"}
    local msg=${2:-""}
    local ts=$(date +"%H:%M:%S.%N" | cut -c-12)

    local reset=$"\e[0m"
    local color=""

    case "${lvl,,}" in
        debug)         color="\e[2m"   ;; # dim
        info)          color="\e[36m"  ;; # cyan
        warn|warning)
            color="\e[33;1m" # yellow bold
            lvl="warn"
        ;;
        error|err)     color="\e[31;1m";; # red bold
        *)             color=""        ;;
    esac

    printf "%b%s [%s] %s%b\n" "$color" "$ts" "$lvl" "$msg" "$reset"
}

section_start "ci_env_section" "CI environment variables" 1
section_start "ci_env_section" "CI environment variables" true
section_start "ci_env_section" "CI environment variables" yes
section_start "ci_env_section" "CI environment variables" on

section_start "ci_env_section" "CI environment variables" off
section_start "ci_env_section" "CI environment variables" false

log_line warn "Сборка началась"
log_line info "Сборка началась"
log_line "warning" "Осталось мало места"
log_line "error" "Не удалось подключиться"
log_line "debug" "Детали шага 1"
log_line "" "Сборка началась"

section_end "ci_env_section"

