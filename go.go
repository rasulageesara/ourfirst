package main

import "fmt"

var num1 float32
var num2 float32
var output float32
var operator string

func main(){

fmt.Print("Enter No 1:")	
fmt.Scanln(&num1)
fmt.Print("Enter No 2:")	
fmt.Scanln(&num2)
fmt.Print("Enter operator:")	
fmt.Scanln(&operator)

if (operator=="+"){
	fmt.Print("your answer is ")
	fmt.Println(add(num1,num2))
	fmt.Println("\n")
}

if (operator=="-"){
	fmt.Print("your answer is ")
	fmt.Println(substract(num1,num2))
	fmt.Println("\n")
}

if (operator=="*"){
	fmt.Print("your answer is ")
	fmt.Println(multiplication(num1,num2))
	fmt.Println("\n")
}


if (operator=="/"){
	fmt.Print("your answer is ")
	fmt.Println(division(num1,num2))
	fmt.Println("\n")
}



}
func add(a float32 ,b float32) float32{
	return a+b
}
func substract(a float32 ,b float32) float32{
   return a-b
}
func multiplication(a float32 ,b float32) float32{
	return a*b
 }
 func division(a float32 ,b float32) float32{
	return a/b
 }