# context shutdown: defer cancel vs wait

`defer cancel()` шлёт сигнал остановки, но не ждёт cleanup worker'а.
Когда `main` выходит, процесс завершается - горутина может не успеть.

## Запуск

```bash
go run . -mode bad    # часто нет "cleanup done"
go run . -mode good   # всегда есть "cleanup done"
go run .              # оба подряд
```

## Суть

| | bad | good |
|---|-----|------|
| cancel | defer при exit main | явный cancel |
| wait | нет | `<-done` или `wg.Wait` |
| cleanup | обрывается с процессом | успевает завершиться |
