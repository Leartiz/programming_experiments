
__timestamp_only_time() {
    echo $(date +"%H:%M:%S.%N" | cut -c-12)
}

__timestamp() {
    echo $(date +%s)
}

# ------------------------------------------------------------------------

section_start() {
  local name=$1
  echo "[===] section_start: $(__timestamp_only_time): ${name} [===]"
  echo ">>> ${name}"
}

section_end() {
  local name=$1
  echo "[===] section_end: $(__timestamp_only_time): ${name} [===]"
  echo
}

log_info() {
    echo "[INFO] $*";
}

log_warn() {
    echo "[WARN] $*";
}

# ------------------------------------------------------------------------

declare -A __T_START

timer_start() {
  local name=$1
  __T_START["$name"]=$(__timestamp)
  section_start "$name"
}

timer_end() {
    local name=$1
    local start=${__T_START[$name]:-}
    local end=$(__timestamp)

    if [[ -n $start ]]; then
        local dur=$(( end - start ))
        log_info "step=$name duration_sec=$dur"
        unset __T_START["$name"]
    else
        log_warn "timer_end called for '$name' without timer_start"
    fi
    section_end "$name"
}

timer_start "Install deps"
sleep 2s
timer_end "Install deps"

section_start "Lint (plain section)"
sleep 1s
section_end "Lint (plain section)"

timer_start "Build"
sleep 0.25s
timer_end "Build"