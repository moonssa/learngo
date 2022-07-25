package main

import "fmt"

type  person struct {
	name string
	age int
	favFood []string
}


func main() {
	favFood :=[]string{"kimchi", "ramen"}
	moon := person {"moon", 25, favFood} // 필드가 의미하는 바가 명확하지 않다. 아래와 같이 사용하는게 좋다
	heon := person {name:"heon", age:26, favFood:favFood}


	fmt.Println(moon)
	fmt.Println(heon)

	a :=2
	b := &a
	*b = 20
	fmt.Println(&a, b)
	fmt.Println(*b, a)

	names := [5] string {"bb", "dd", "aa", "hh", "jj"}
	fmt.Println(names)

	nickname :=[] string {"haha"}
	nickname = append(nickname, "hihi")
	fmt.Println(nickname)

	nico :=map[string]string {"name": "nico", "age":"15"}
	fmt.Println(nico)

	for key, value := range nico {
		fmt.Println(key, value)
	}
	for _, value := range nico {
		fmt.Println( value)
	}
}