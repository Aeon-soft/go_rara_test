package main

import (
	"fmt"
	"reflect"
)

type rectangel struct{
	len float64
	bre float64
	color string
	geomatry struct{
		area int 
		perimeter int
	}
}

type Salary struct{
	date string
	amount int
}
type Employee struct{
	FirstName , LastName , Email string
	Age int 
	salary []Salary 
}
func main(){
	var	r rectangel
	r.len=12
	r.bre =2
	r.color = "oo"
	r.geomatry.area=12
	r.geomatry.perimeter = 33
	fmt.Println("hello struct type",r)
	v := reflect.ValueOf(r)
	fmt.Println(v)

	s1 := Salary{"2020",10000}
	s2 := Salary{"2020",10000}
	s3 := Salary{"2020",10000}
	s4 := Salary{"2020",10000}
	e1 := Employee{FirstName:"pramish",LastName:"karkee",Email:"karkipramish18@gmail.com",Age:22,salary:[]Salary{s1,s2,s3,s4}}
	fmt.Println(e1)
	var arr [] Salary
	arr = append(arr,s1)
	arr = append(arr,s2)
	fmt.Println(arr)
}