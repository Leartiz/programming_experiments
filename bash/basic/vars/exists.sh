#!/bin/dash

SC_COMMON_ENV_VAR=1

if [ "${SC_COMMON_ENV_VAR+set}" = set ]; then
    echo "SC_COMMON_ENV_VAR существует"
else
    echo "SC_COMMON_ENV_VAR не существует"
fi

define_if_not_exists() {
    var_name="$1"
    default_value="$2"

    value="${default_value}" # remove warnings
    eval "value=\${${var_name}+set}"
    if [ "$value" != "set" ]; then
        eval "$var_name=\"\$default_value\""
        eval "export $var_name"
    fi
}

echo "${SC_COMMON_LOGGING_CREATE_ALIAS_WITHOUT_NS}"
define_if_not_exists "SC_COMMON_LOGGING_CREATE_ALIAS_WITHOUT_NS" ""
echo "${SC_COMMON_LOGGING_CREATE_ALIAS_WITHOUT_NS}"
define_if_not_exists "SC_COMMON_LOGGING_CREATE_ALIAS_WITHOUT_NS" "123"
echo "${SC_COMMON_LOGGING_CREATE_ALIAS_WITHOUT_NS}"
