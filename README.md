# go-gb
Golang practice

##HW-6
1. Проанализируйте задания предыдущих уроков.
   a. Q: В каких случаях необходима была явная передача указателя в качестве входных параметров и возвращаемых результатов или в качестве приёмника в методах?
   A: На примере предыдущих уроков, передача указателя в качестве параметров была нужна в случаях когда рекурсивно работать с мапой (в предыдущем задании), где  
   переход от передачи по значению (через копирование объекта) к передаче по ссылке (через поинтер) позволил еще в несколько раз ускорить оптимизированный подсчет чисел Фибоначчи через рекурсию, результат по скорости приближается к наиболее оптимальному - подсчет в памяти, без рекурсии.
   Результат следующий:
   Фибоначчи 45-е число, методом обычной рекурсии: 7.117149724s
   Фибоначчи 45-е число, методом рекурсии с кешированием в мапе: 23.74µs
   Фибоначчи 45-е число, методом рекурсии с кешированием в мапе + передача по ссылке: 7.396µs
   b. Q: В каких случаях мы фактически имеем дело с указателями при передаче параметров, хотя явно их не указываем?
   A: Например, в случае использования defer, значения захватываются по указателю при передаче в функции внутри блока. Когда передаем в параметре слайс или мапу.

2. Q: Для арифметического умножения и разыменования указателей в Go используется один и тот же символ — оператор (*). Как вы думаете, как компилятор Go понимает, в каких случаях в выражении имеется в виду умножение, а в каких — разыменование указателя?
   A: При использовании арифметического умножения компилятор ожидает значения двух переменных, оператор разыменывания же является префиксом единственной переменной. Соответственно если переменная одна, значит * это префикс и оператор разыменывания, иначе умножение.


