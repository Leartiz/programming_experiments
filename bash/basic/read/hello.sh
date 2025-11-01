log() {
    echo "$*"

    local tm lvl msg

    # NOTE:
    #  -r	do not allow backslashes to escape any characters
    #
    read -r tm lvl msg <<< "$*"

    echo $tm
    echo $lvl
    echo $msg
}

log "$(date +"%H:%M:%S.%N" | cut -c-12) info msg msg msg"