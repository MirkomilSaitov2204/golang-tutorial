package main

import "fmt"

func main(){
	var firstName string = "Mirkomil"
	var lastName string = "Saitov"
	var age int = 30
	var height float64 = 1.75
	var isStudent bool = false

	fmt.Println("Ism:", firstName)
	fmt.Println("Familiya:", lastName)
	fmt.Println("Yosh:", age)
	fmt.Println("Bo'y:", height, "metr")
	fmt.Println("Talaba emasmi?", isStudent)

	firstName = "Aziz"
	age = 25
	fmt.Println("Yangi ism:", firstName)
	fmt.Println("Yangi yosh:", age)

	const class = "Go Programming"
	fmt.Println("Kurs:", class)


}