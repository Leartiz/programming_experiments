__timestamp_only_time() {
    echo $(date +"%H:%M:%S.%N" | cut -c-12)
}

__log_to_console() {
    local json_output=$1
    printf "%s\n" "$json_output"
}

__LOG_DIR="./logs"
__LOG_JSON_FILE="log.jsonl"
__log_to_file() {
    local json_output=$1
    mkdir -p "$__LOG_DIR"
    printf "%s\n" "$json_output" >> "${__LOG_DIR}/$__LOG_JSON_FILE"
}

# ------------------------------------------------------------------------

logging::create_log_entry() {
    local level=${1:-info}
    local message=${2:-}

    log_entry["ts"]=$(__timestamp_only_time)
    log_entry["level"]="$level"
    log_entry["project"]="${CI_PROJECT_PATH:-local}"
    log_entry["pipeline"]="${CI_PIPELINE_ID:-na}"
    log_entry["job"]="${CI_JOB_ID:-na}"
    log_entry["sha"]="${CI_COMMIT_SHA:-na}"
    log_entry["message"]="$message"
}

logging::log_to_json() {
    local level="$1"
    local message="$2"

    # NOTE:
    #   вернуть "массив" из функции корректно не получилось
    #   вариант объявить и вызвать создание (рабочий).
    #     Причём есть возможность переопределить!
    #
    declare -A log_entry
    logging::create_log_entry "${level}" "${message}"

    local json_output='{'
    local first=true

    for key in "${!log_entry[@]}"; do
        if [[ $first == true ]]; then
            first=false
        else
            json_output+=','
        fi
        json_output+="\"$key\":\"${log_entry[$key]}\""
    done

    json_output+='}'

    __log_to_console "${json_output}"
    __log_to_file "${json_output}"
}

# manual test
# ------------------------------------------------------------------------

logging::log_to_json "trace" "test log string"
logging::log_to_json "debug" "test log string"
logging::log_to_json "info" "test log string"
logging::log_to_json "warning" "test log string"
logging::log_to_json "error" "test log string"
logging::log_to_json "fatal" "test log string"
