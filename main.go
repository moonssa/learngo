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

func main(){
	totalLen, Upp := lenAndUpper("moon")
	fmt.Println(totalLen, Upp)
	tLen , _ := lenAndUpper("moon") // return 값 생략 가능 _ 이용
	fmt.Println(tLen)
	repeatMe("moon", "happy", "thanks")
	
	totalLen2, Upp2:= lenAndUpper2("honey")
	fmt.Println(totalLen2, Upp2)

}