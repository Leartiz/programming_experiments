__timestamp() {
    echo $(date +%s)
}

log_info() {
    echo "[INFO] $*";
}

log_warn() {
    echo "[WARN] $*";
}

# ------------------------------------------------------------------------

declare -A __T_START
name="timer"

__T_START["$name"]=$(__timestamp)
echo "__T_START["$name"]: ${__T_START["$name"]}"

start=${__T_START[$name]:-}
echo "start: $start"

sleep 2.5s

end=$(__timestamp)
echo "end: $end"

if [[ -n $start ]]; then
    dur=$(( end - start ))
    log_info "step=$name duration_sec=$dur"
    unset __T_START["$name"]
else
    log_warn "timer_end called for '$name' without timer_start"
fi

# NOTE:
#   значение было сброшено!
#
echo "__T_START["$name"]: ${__T_START["$name"]}"