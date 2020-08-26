// Указатели project main.go
//* int - Указатель на int
package main

import (
	"fmt"
	"reflect" //reflect.Type - выводим типы указателей
)

func main() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt)) //Получаем указатель на myInt и выводим тип указателя
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat)) //Получаем указатель на myInt и выводим тип указателя
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool)) //Получаем указатель на myInt и выводим тип указателя

	//Переменные содержащие указатели
	/*var myInt int
	var myIntPointer *int // Объявление переменной содержащий указатель на int
	myIntPointer = &myInt //Указатель присваивается переменной
	fmt.Println(myIntPointer)*/

}
