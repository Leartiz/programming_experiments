echo_color() {
    local message=${1:-""}
    local color_code=${2:-255}
    echo -e "\e[38;5;${color_code}m"$message"\e[0m"
}

for ((color_code=0; color_code < 256; color_code++))
do
    echo_color "test color" "$color_code"
done

