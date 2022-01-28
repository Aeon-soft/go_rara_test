package main

import (
	"fmt"
)
func changeValue(str *int){
	*str = 60
}

func changeValue2(str int){
	str = 60
}

type Point struct {
	x int32
	y int32
}

type Circle struct {
	radious float64
	center *Point
}
func UpdateValue(p *Point){
	p.x =200
}
func main(){
	var p1 Point = Point{x:2,y:40}
	var p2 Point = Point{3,-1}
	var a=20
	p3 := &Point{x:3}
	fmt.Println(p1.x,p2,p3,a)
	UpdateValue(&p1)
	fmt.Println(p1.x,p2,p3,a)
	fmt.Println("circle")
	var c1 Circle = Circle{4.1, &Point{x:2,y:40}}
	fmt.Println(c1,c1.center.x  )
	toChange := 40
	fmt.Println(toChange)
	// changeValue(&toChange)
	fmt.Println(toChange)
	fmt.Println(*&toChange,toChange)
}