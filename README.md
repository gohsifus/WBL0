# Запуск
NatsStreaming в docker контейнере
```
make runNatsStreaming
```
Запуск сервиса
```
make runService
```
Запуск publisher для публикации данных
```
make runPub
```
Тесты
```
make test
```
# Валидация данных из канала
Перед записью при декодировании json проверяется его структура, если в json отсутсвуют поля помеченные как required данные игнорируются, аналогично в случае если из канала получен не json.
# Хранение данных
Структуру таблиц и модели для работы с ними можно посмотреть в migrations и pkg/models.
# Публикация данных
В cmd/pub функционал для генерации и публикации в канал случайных данных. 
# Сохранность данных
Для того чтобы не терять данные в случае падения сервиса используется durable подписка.
# Стресс тест
![Alt text](https://github.com/gohsifus/WBL0/blob/dev/images/vegetaRep.png "Report") 
График
![Alt text](https://github.com/gohsifus/WBL0/blob/dev/images/vegetaPlot.png "Graph") 
Кеширование шаблона позволило увеличить количество запросов/c. в 3 раза.
