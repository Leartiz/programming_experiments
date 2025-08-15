# NOTE:
#   разные способы создать файл
#
touch newfile_0.txt
> newfile_1.txt

OUT_FILE_NAME=out.txt

# NOTE:
#   > (создать/перезаписать)
#
printf "Начало лога\n" > ${OUT_FILE_NAME}

{
    for i in {1..3}; do
        printf "Строка номер %d\n" "$i"
    done
} >> ${OUT_FILE_NAME}

# NOTE:
#   >> (дописать)
#
printf "Конец лога\n" >> ${OUT_FILE_NAME}

# ------------------------------------------------------------------------

OUT_FILE_NAME=logs.jsonl

logs='[
    {"ts":"2025-08-15T10:00:00Z","level":"info","msg":"start"},
    {"ts":"2025-08-15T10:00:05Z","level":"warn","msg":"slow request"},
    {"ts":"2025-08-15T10:00:10Z","level":"error","msg":"db timeout"}
]'

> "$OUT_FILE_NAME"

# NOTE:
#   -c (compact, компактный вывод без лишних пробелов)
#   входной JSON - массив, то '.[]' выдаст каждый элемент массива как отдельный JSON-объект
#
jq '.[]' <<< "$logs" >> "$OUT_FILE_NAME"