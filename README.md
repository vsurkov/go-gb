# go-gb
Golang practice


##HW-8
В /cmd/ismg_exporter добавлено приложение которое умеет читать свою конфигурацию из ENV, а при их отсутсвии из флагов, если нет ни первого, ни второго - возьмет значение по умолчанию.
Для параметров типа Port и URL добавлена валидация значения - при не валидном значении будет взято значение по умолчанию.


##HW-7
**Q:** 1. С какими интерфейсами мы уже сталкивались в предыдущих уроках? Обратите внимание на уроки, в которых мы читали из 
стандартного ввода и писали в стандартный вывод.
**A:** Мы сталкивались интерфейсами error в фукциях порождающих ошибку, с interface{} в fmt.Scan* и fmt.Print* где это используется
как способ сделать функции универсальными под разные типы.

**Q:** 2. Посмотрите примеры кода в своём портфолио. Везде ли ошибки обрабатываются грамотно? Хотите ли вы переписать какие-либо функции?
**A:** Основные ошибки были с ошибками были исправлены в первых заданиях, после кодревью, но сейчас бы я изменил одно - 
сейчас обработка ошибки происходит в main, более правильным было бы обаботать ее в роутере и попытаться исправить, 
запросив у пользователя повторный ввод корректных данных и уже при невозможности (по какой либо причине) выбрасывать в 
main и там выходить с кодом 1
Что вроде было сделано правильно:
- в месте возникновения ошибки в вывод выводится причина и ошибка выбрасывается выше
- ошибки не обрабатываются в месте возникновения а возвращаются в инициализирующую функцию (наверх)
- в коде нет panic()

##HW-5
### Назначение
Приложение использует инициализированное в main() значение num := 30 для подсчета чисел Фибоначчи двумя разными реализациями

### Структура
Приложение разделено на:
1. **main.go** - реализует общую логику работы приложения:
2. **fibonacci.go** - реализует методы получения подсчета чисел Фибоначчи двумя реализациями:
- Fib(n int) int - обычный способ, через рекурсию
- GetFibonacci(n int) int - улучшенный способ, через рекурсию и хранение значений в map[int]int
3. **fibonacci_test.go** - unit-тесты для обоих методов реализации:
В тесте используется мапа со значениями от 0 до 20, и проверка на соответствие предопределенного значения полученному 

##HW-4
### Назначение
Приложение принимает на вход набор целых чисел и выводит его же в отсортированном виде. На данный момент реализован алгоритм сортировки методом вставки.
Приложение размещено в /pkg/my_sort для дальнейшего использования.

### Структура
Приложение разделено на:
1. **main.go** - точка входа 
2. **my_sort.go** - реализует общую логику работы приложения и последующего расширения логики: 
> добавление альтернативных способов ввода (файл, канал, сокет?);
> добавление альтернативных алгоритмов сортировки;
> добавление альтернативных способов вывода (файл, канал, сокет?);

3. **request.go** - реализует методы получения входных данных;
4. **data_prepare.go** - реализует обработку и подготовку данных;
5. **inserts_sort.go** - реализует алгоритм сортировки вставками
6. **respond.go** - реализует методы вывода данных;
7. **helper.go** - реализует вспомогательные функции внутреннего использования.
