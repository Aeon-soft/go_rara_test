package main 

import (
	"fmt"
)

type Rectangle struct {
	length int
}
func main(){
	var arr [5]int
	var rect Rectangle = Rectangle{11}
	rect.length = 100
	rect.length = 102

	var rect1 Rectangle
	rect1.length=1004
	ChangeMe(&rect)
	fmt.Println("**rect",rect)
	arr2D := [2][2]int {{1,2},{3,4}}
	arr[0] = 100
	fmt.Println(arr)
	str := "im"

	arr1 := [3]int{2,3,4}
	fmt.Println(arr1,len(arr1),len(str))
	sum :=0
	for i := 0; i<len(arr1); i++ {
		sum += arr1[i]
	}
	fmt.Println("sum",sum,arr2D)

	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	fmt.Println(strDict["Japan"],strDict)
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}
	for key := range strDict {
		fmt.Println(key)
	}
	for range "Hello" {
		fmt.Println("Hello")
	}

	func (l int, b int){
		fmt.Println(l*b)
	}(20,30)
}
// https://www.golangprograms.com/go-language/variables.html