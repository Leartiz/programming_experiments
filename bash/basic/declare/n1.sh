try_declare_global() {
    declare -g RESULT
    RESULT="ok"
}

try_declare_int() {
    declare -i -g n=1+2
}

# -----------------------------------------------

try_declare_global
echo "RESULT: $RESULT"

try_declare_int
echo "n: $n"