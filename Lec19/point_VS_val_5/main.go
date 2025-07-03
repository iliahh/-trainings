package main

import "fmt"

type Rectangle struct {
	length, width int
}

//1. Реализуем функцию и метод для подсчёта периметра прямоугольника
// ВАЖНО: Все параметры передаём как копии

//4. При вызове этого метода не важно, будет ли он вызываться у экземпляра или у его ссылки
func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

//5. Данную функцию можно вызывать ТОЛЬКО у копии прямоугольника (но не у его ссылки)
func Perimeter(r Rectangle) int {
	return 2 * (r.length + r.width)
}

//6. Допустим будет метод, который меняет значение состояния структуры на новое, но этот метод - Value Reciever
func (r Rectangle) SetLength(newLength int) {
	r.length = newLength
}

func main() {
	//2. Создаём экземпляр прямоугольника
	rectangleAsValue := Rectangle{10, 10}
	fmt.Println("Call func for rectangleAsValue:", Perimeter(rectangleAsValue))    //Call func for rectangleAsValue: 40
	fmt.Println("Call method for rectangleAsValue:", rectangleAsValue.Perimeter()) //Call method for rectangleAsValue: 40

	//3. Создадим указатель на прямоугольник
	rectangleAsPointer := &rectangleAsValue
	fmt.Println("Call method for rectangleAsPointer:", rectangleAsPointer.Perimeter()) //Call method for rectangleAsPointer: 40
	//Perimeter(rectangleAsPointer) - иллюстрация к пункту 5

	//7. Вызываем метод SetLength у экземпляра rectangleAsValue
	fmt.Println("Before call method SetLength:", rectangleAsValue) //Before call method SetLength: {10 10}
	rectangleAsValue.SetLength(9999)
	fmt.Println("After call method SetLength:", rectangleAsValue) //After call method SetLength: {10 10}

	//8. Вызываем метод SetLength у ссылки на rectangleAsValue
	rectangleAsPointer.SetLength(999999999)
	fmt.Println("After call method SetLength for &rectangle:", *rectangleAsPointer) //After call method SetLength for &rectangle: {10 10}

	//Ничего не меняется, потому что внезависимости от того, передаём ли мы в метод значение по ссылке
	//  или передаём его как value, если в сигнатуре метода в принципе указано, что он принимает значения
	//  как value, то он и будет все действия над нашим экземпляром совершать как над value.
	//Не смотря даже на то, что у метода на 23 строке происходят какие-то изменения нашего объекта,
	//  при вызове этого метода на 44 строке у указателя на прямоугольник происходит следующее:
	//  мы разыменовываем указатель при вызове, отрабатываем с локальной копией, и после этого возвращаем
	//  управление вызывающей стороне. Поскольку мы отработали с локальной копией,
	// естественно у исходного прямоугольника никаких изменений не произошло.
	//Это очень классная штука, поскольку нужно помнить только про сам метод, а на этапе его реализации
	//  можно заложить любую логику и используя написание valueReciever'а мы можем гарантировать,
	//  что на вызывающей стороне никаких изменений с нашим объектом не произойдет.
	//Но с другой стороны, если использовать в такой ситуации valueReciever, не смотря на то,
	//  что мы тут вызываем метод у pointer'а, он будет разыменован, и метод будет работать с копией,
	//  то есть мы скопируем всё равно в этой ситуации всё, что было под этим адресом.
	//  Здесь решает в общем сама сигнатура метода.
}
