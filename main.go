package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string)(int, string){
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper2(name string)(length int, uppercase string) {
	defer fmt.Println("I'm done")
	fmt.Println("defer test")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return  // naked return
}

func superAdd_test(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	fmt.Println("traditional")
	for i:=0; i < len(numbers); i ++ {
		fmt.Println(numbers[i])
	}
	return 1
}
func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		fmt.Println(number)
		total += number
	}
	return total
}
func canIDrink(age int) bool {
	if koreanAge := age+2; koreanAge < 18 {
		return false
	} 
	return true
	
}

func canIDrinkTest(age int) bool {
	switch koreanAge := age+2 ; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}
func main(){
	fmt.Println(canIDrinkTest(8))
	fmt.Println(canIDrinkTest(16))

	// fmt.Println(canIDrink(15))
	// fmt.Println(canIDrink(16))

	// result := superAdd(1,3,5,7,9,11)
	// fmt.Println(result)
	// totalLen, Upp := lenAndUpper("moon")
	// fmt.Println(totalLen, Upp)
	// tLen , _ := lenAndUpper("moon") // return 값 생략 가능 _ 이용
	// fmt.Println(tLen)
	// repeatMe("moon", "happy", "thanks")
	
	// totalLen2, Upp2:= lenAndUpper2("honey")
	// fmt.Println(totalLen2, Upp2)

}