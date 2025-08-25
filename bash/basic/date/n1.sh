#!/bin/bash

LOGGING_TIMESTAMP_FMT="iso8601"

timestamp_fmt() {
    local default_value=${LOGGING_TIMESTAMP_FMT:-"unix"}
    local fmt="${1:-$default_value}"

    case "${fmt,,}" in
        unix)      date +%s ;;
        unix_ms)   date +%s%3N ;;
        time)      date +"%H:%M:%S" ;;
        time_ms)   date +"%H:%M:%S.%N" | cut -c-12 ;;
        iso8601|*) date +"%Y-%m-%dT%H:%M:%S%z" ;;
    esac
}

# test default timestamp_fmt
# ------------------------------------------------------------------------

timestamp_fmt unix
timestamp_fmt
timestamp_fmt
echo

# test flag -u
# ------------------------------------------------------------------------

echo "date -u +%s:        $(date -u +%s)"
echo "date +%s:           $(date +%s)"
echo "timestamp_fmt unix: $(timestamp_fmt unix)"
echo

echo "date -u +%s%3N:        $(date -u +%s%3N)"
echo "date +%s%3N:           $(date +%s%3N)"
echo "timestamp_fmt unix_ms: $(timestamp_fmt unix_ms)"
echo

echo "date -u +"%H:%M:%S":   $(date -u +"%H:%M:%S")"
echo "date +"%H:%M:%S":      $(date +"%H:%M:%S")"
echo "timestamp_fmt time:    $(timestamp_fmt time)"
echo "timestamp_fmt time_ms: $(timestamp_fmt time_ms)"
echo

echo "date -u +"%Y-%m-%dT%H:%M:%S%z": $(date -u +"%Y-%m-%dT%H:%M:%S%z")"
echo "date +"%Y-%m-%dT%H:%M:%S%z":    $(date +"%Y-%m-%dT%H:%M:%S%z")"
echo "timestamp_fmt iso8601:          $(timestamp_fmt iso8601)"
echo

# ------------------------------------------------------------------------