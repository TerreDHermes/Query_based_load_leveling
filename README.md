# Query_based_load_leveling

## Описание данного паттерна

  Query-based Load Leveling — паттерн интеграции между приложениями на основе очередей с целью балансировки нагрузки. Использует очередь для сглаживания нагрузки на сервис. 

  Периодические тяжелые нагрузки могут «ронять» сервис или повышать тайм-аут обработки задачи. Данный паттерн сглаживает влияние пиков на доступность и оперативность реагирования. Пример — использование очереди на входе сервиса сбора логов.

  Как работает: у вас есть сервис, и есть очередь сообщений, куда падают таски. Таски могут генерироваться любой системой. Вы знаете, что в какой-то момент времени есть всплески по задачам. 

  Нагрузочное тестирование показало, что ваша конфигурация выдерживает 100 запросов. Но вы работаете по HTTP и понимаете, что у вас всплеск — 150 запросов. По идее, в такой момент сервис может что-то отложить и обработать позже, а может свалиться от количества запросов и перестать отвечать.

  Чтобы избежать этого, мы ставим очередь на вход — если раньше был запрос напрямую к сервису, то теперь мы всё делаем на очередях. Таски падают в очередь, и сервис их обрабатывает. 

  ## Текущая реализация 

  Создается 10 горутин на уровне слоя сервиса, запросы складываются в буферизированный канал. 10 горутин обрабатывают очередь из запросов (запросы, которые лежат в канале).
