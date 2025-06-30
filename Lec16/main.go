package main

import "fmt"

//1. Указатели на массивы. Почему так делать не надо
func mutation(arr *[3]int) {
	// (*arr)[1] = 909
	// (*arr)[2] = 100000
	//Можно написать и так, т.к. Go сам разыменует указатель в массив (из-за того, что функция принимает *arr)
}

//2. Используйте лучше слайсы (это идеоматично с точки зрения Go) (Слайсы уже обладают нужными указателями, поэтому лучше не изобретать велосипед)
func mutationSlice(sls []int) {
	sls[1] = 909
	sls[2] = 10000
}

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println("Arr before mutation:", arr) //Arr before mutation: [1 2 3]
	mutation(&arr)
	fmt.Println("Arr after mutation:", arr) //Arr after mutation: [1 909 100000]

	newArr := [3]int{1, 2, 4}
	fmt.Println("newArr before mutationSlice:", newArr) //newArr before mutationSlice: [1 2 4]
	mutationSlice(newArr[:])
	fmt.Println("newArr after mutataionSlice:", newArr) //newArr after mutataionSlice: [1 909 10000]
}
