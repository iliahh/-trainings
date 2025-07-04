🔸 1. Логические операторы (boolean)
Работают с типом bool (true или false):

Оператор	Название	Описание	Пример
&&	И (AND)	Истина, если обе части истины	true && true → true
`		`	ИЛИ (OR)
!	НЕ (NOT)	Инвертирует значение (меняет true на false и наоборот)	!true → false

📌 Пример:

x := 5
fmt.Println(x > 3 && x < 10) // true
fmt.Println(x < 3 || x == 5) // true
fmt.Println(!(x == 5))       // false

🔸 2. Операторы сравнения
Сравнивают значения. Возвращают true или false.

Оператор	Название	Пример	Результат
==	Равно	5 == 5	true
!=	Не равно	5 != 3	true
<	Меньше	3 < 5	true
>	Больше	7 > 2	true
<=	Меньше или равно	5 <= 5	true
>=	Больше или равно	6 >= 3	true

🔸 3. Арифметические операторы
Работают с числами (int, float64, и т.д.):

Оператор	Описание	Пример	Результат
+	Сложение	5 + 3	8
-	Вычитание	5 - 2	3
*	Умножение	2 * 4	8
/	Деление	10 / 2	5
%	Остаток от деления	7 % 3	1

🔸 4. Операторы присваивания
Используются, чтобы записать значение в переменную:

Оператор	Пример	То же, что и
=	x = 5	просто присваивание
+=	x += 3	x = x + 3
-=	x -= 2	x = x - 2
*=	x *= 4	x = x * 4
/=	x /= 2	x = x / 2
%=	x %= 3	x = x % 3

🔸 5. Инкремент и декремент (но Go больше не разрешает их в выражениях)
Оператор	Название	Описание
x++	Инкремент	Увеличить x на 1
x--	Декремент	Уменьшить x на 1

⚠️ В Go нельзя использовать x = x++ + 2, как в других языках. x++ можно писать только отдельно как инструкцию.

🔸 6. Операторы работы с битами (для продвинутого уровня)
Оператор	Название	Пример	Описание
&	И (AND) побитовый	a & b	1, если биты в обоих — 1
`	`	ИЛИ (OR) побитовый	`a
^	XOR (исключающее ИЛИ)	a ^ b	1, если биты разные
&^	AND NOT	a &^ b	Очищает биты a, где b — 1
<<	Сдвиг влево	a << 2	Умножение на 2²
>>	Сдвиг вправо	a >> 1	Деление на 2¹

✅ Заключение
Если ты новичок, то сначала тебе чаще всего понадобятся:

&&, ||, ! — логические операторы

==, !=, <, > — сравнение

+, -, *, /, % — арифметика

= и +=, -= и т.д. — для обновления значений
