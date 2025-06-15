package main

import "fmt"

func main() {
	//Switch!
	/*
		if price == 100 {
			fmt.Println("First case")
		} else if price == 110 {
			fmt.Println("Second case")
		} else if price == 120 {
			fmt.Println("Third case")
		}.....
	*/
	price := 120
	fmt.Scan(&price)
	//В switch-case запрещены дублирующиеся условия в case'ах
	switch price {
	case 100:
		fmt.Println("First case")
	case 110:
		fmt.Println("Second case")
	case 120:
		fmt.Println("Third case")
	case 130:
		fmt.Println("Another case")
	default:
		//Отрабатывает только в том случае, если не один из выше перечисленных кейсов не сработал
		fmt.Println("Default case")
	}

	//Case с множеством вариантов
	var ageGroup string = "q" //Возрастные группы : "a", "b", "c", "d", "e"
	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("AgeGroup 10-40")
	case "d", "e":
		fmt.Println("AgeGroup 50-70")
	default:
		fmt.Println("You are too young/old")
	}

	//Case с выражениями
	var age int
	fmt.Scan(&age)

	switch {
	case age <= 18:
		fmt.Println("Too young")
	case age > 18 && age <= 30:
		fmt.Println("Second case")
	case age > 30 && age <= 100:
		fmt.Println("Too old")
	default:
		fmt.Println("Who are you")
	}

	//Case с проваливаниями. Проваливания выполняют ДАЖЕ ЛОЖНЫЕ КЕЙСЫ
	//В момент выполнения fallthrough у следующего кейса не проверяется условие,
	//а сразу выполняется тело.
	//continue не работает в свиче, в отличие от break
	var number int
	fmt.Scan(&number)
	//outer:
	switch {
	case number < 100:
		fmt.Printf("%d is less then 100\n", number)
		if number%2 == 0 {
			break //outer
			//с брейком код отработает только в этом теле блока,
			// то есть выдаст "7 is less then 100"
			//Можно добавить лейбл, но их использовать не рекомендуется
		}
		fallthrough //команда, которая позволяет провалиться дальше по кейсам,
		// ведь без нее при введении числа 7 выдаст только "7 is less then 100",
		// а с этой командой также выдаст с прошлой строкой "d7 less then 200"
	case number < 200:
		fmt.Printf("%d less then 200\n", number)
		fallthrough //если бы тут не было фолтру,
		// тогда бы код не прошёл до значения тысячи и вывел бы лишь 2 строки с <100 и <200
	case number > 1000:
		fmt.Printf("%d GREATER then 1000\n", number)
		fallthrough
	default:
		fmt.Printf("%d DEFAULT\n", number)
	}

	//Гадость с терминацией цикла for из switch
	var i int
uberloop:
	for {
		fmt.Scan(&i)
		switch {
		case i%2 == 0:
			fmt.Printf("Value is %d and it's even\n", i)
			break uberloop
		}
	}
	fmt.Println("END")

	//ЕСЛИ ВЫ ИСПОЛЬЗУЕТЕ ВЛОЖЕННЫЕ ЦИКЛЫ - ИСПОЛЬЗУЙТЕ ОБЯЗАТЕЛЬНО ЛЕЙБЛЫ!
	//ЕСЛИ ВЫ ИСПОЛЬЗУЕТЕ switch-case'ы ВЛОЖЕННЫЕ В ЦИКЛЫ - ОБЯЗАТЕЛЬНО ИСПОЛЬЗУЙТЕ ЛЕЙБЛЫ!
	//ЕСЛИ ИСПОЛЬЗУЕТЕ ОДИНОКО СТОЯЩИЙ for ИЛИ switch - НЕ ИСПОЛЬЗУЙТЕ ЛЕЙБЛЫ!
	//Если к примеру три цикла друг в друге, то нужно 2 или лучше 3 лейбла
	//break как бы "зашит" в каждый кейс и как раз fallthrough позволяет его избежать
}
