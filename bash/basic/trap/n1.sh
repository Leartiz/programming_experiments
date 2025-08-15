# NOTE:
#   https://tokmakov.msk.ru/blog/item/96

count=0

trap 'echo "Получен сигнал SIGINT \
(нажатие Ctrl+C), завершение сценария"; exit 1' SIGINT

while true ; do
    sleep 0.25s
    (( count++ ))
    echo $count
done

