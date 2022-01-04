# go-gb
Golang practice

##HW-5
### Назначение
Приложение использует инициализированное в main() значение num := 30 для подсчета чисел Фибоначчи двумя разными реализациями

### Структура
Приложение разделено на:
1. **main.go** - реализует общую логику работы приложения:
2. **fibonacci.go** - реализует методы получения подсчета чисел Фибоначчи двумя реализациями:
- Fib(n int) int - обычный способ, через рекурсию
- FibM(n int) int - улучшенный способ, через рекурсию и хранение значений в map[int]int
3. **fibonacci_test.go** - unit-тесты для обоих методов реализации:
   В тесте используется мапа со значениями от 0 до 20, и проверка на соответствие предопределенного значения полученному 