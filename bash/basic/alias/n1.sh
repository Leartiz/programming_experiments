__timestamp() {
    echo $(date -u +"%Y-%m-%dT%H:%M:%S%z")
}

__timestamp_

__context() {
    echo "proj=${CI_PROJECT_PATH:-local} \
pipe=${CI_PIPELINE_ID:-na} \
job=${CI_JOB_ID:-na} \
sha=${CI_COMMIT_SHORT_SHA:-na}"
}

echo "context: $(__context)"
echo "timestamp: $(__timestamp)"