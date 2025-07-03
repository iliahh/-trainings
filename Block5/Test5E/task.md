E. Go в отрезки . V2
Ограничение времени	1 секунда
Ограничение памяти	64Mb
Ввод	стандартный ввод или input.txt
Вывод	стандартный вывод или output.txt
Виктор - перешел на второй курс. Лекции по аналитической геометрии продолжаются и Виктор решает повторить задания из первого курса, которые напрямую касаются великого вопроса - принадлежит ли точка отрезку или не принадлежит? Но теперь Виктор решает задачи , поддерживая ООП. Помогите Виктору реализовать решение задачи используя структуры 
P
o
i
n
t
 и 
L
i
n
e
:
type Point struct { 
        x float64 
        y float64 
} 
 
func NewPoint(newX, newY float64) ∗Point { 
        return &Point{newX, newY} 
} 
 
//Return X coord of Point 
func (p ∗Point) GetX() float64 { 
} 
 
//Return Y coord of Point 
func (p ∗Point) GetY() float64 { 
} 
 
//Set new X coord for Existing point 
func (p ∗Point) SetX(newX float64) { 
} 
 
//Set new Y coord for Existing point 
func (p ∗Point) SetY(newY float64) { 
} 
 
//Retunrt string info about point in format Point{X: value, Y: value} 
func (p ∗Point) Stringify() string { 
} 
 
type Line struct { 
        p1 Point 
        p2 Point 
} 
 
func NewLine(newP1, newP2 Point) ∗Line { 
        return &Line{newP1, newP2} 
} 
 
//Return p1 of line 
func (l ∗Line) GetP1() Point { 
} 
 
//Return p2 of line 
func (l ∗Line) GetP2() Point { 
} 
 
//Set new p1 for Existing Line 
func (l ∗Line) SetP1(newP1 Point) { 
} 
 
//Set new p2 for Existing Line 
func (l ∗Line) SetP2(newP2 Point) { 
} 
 
//Calc length of segment p1p2 
func (l ∗Line) Length() float64 { 
} 
 
//Method return True if p3 belongs to p1p2 segment 
//                       False otherwise 
func (l ∗Line) IsOnSegment(p3 Point) bool { 
} 
 
//Retunrt string info about line in format Line{P1: value, P2: value} 
func (l ∗Line) Stringify() string { 
}
Формат ввода
На вход программе подается 4 целых числа (все в одной строке, разделены пробелами) (
x
1
,
y
1
,
x
2
,
y
2
), а затем с новой строки еще 2 числа (
x
3
,
y
3
).
Формат вывода
На выходе ваша программе выводит 
Y
E
S
 , если точка с координатами (
x
3
,
y
3
) принадлежит отрезку (
x
1
,
y
1
|
x
2
,
y
2
), и 
N
O
 в противном случае.
Пример 1
Ввод
Вывод
0 0 10 10
5 5
YES
Пример 2
Ввод
Вывод
1 1 3 2
1 1
YES
Примечания
Ваше решение должно содержать структуры Point , Line и блоки считывания и вывода данных в консоль / файл. Да, это не опечатка, что на вход программе идут целые числа , а конструктор принимает вещественные параметры. В такой постановке вам будет легче проводить расчетно-вычислительные операции.