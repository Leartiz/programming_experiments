## n42 - `QListIterator; QMutableListIterator; foreach`

- implementation [here](./main.cpp)

## Теория

`Qt` делает копию контейнера при входе в цикл `foreach`,
поэтому если вы будете менять значение элементов в цикле, 
то на оригинальном контейнере это никак не отразится.