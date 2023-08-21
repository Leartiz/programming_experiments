## n33 - sending events

## Теория

Есть два метода для отправки событий: 
`QCoreApplication::postEvent()` и `QCoreApplication::sendEvent()`.

Отправка событий методом `postEvent()` обладает надежностью в потоках, 
а методом `sendEvent()` - нет. Поэтому при работе с разными
потоками всегда используйте метод `postEvent()`.