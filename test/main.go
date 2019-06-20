package main

import "fmt"

func main() {
	defer func() {
		if err := recover();err!=nil{
			fmt.Println("some thing wrong")
		}
	}()
	var name string
	for {
		fmt.Scan(&name)
		if name == "exit"{
			panic("wrong")
		}
		go MyTestForPanic(name)
		fmt.Println(name)
	}
}

func MyTestForPanic(str string)  {
	defer func() {
		if err := recover();err!=nil{
			fmt.Println("some thing wrong")
		}
	}()
	if str == "wukong"{
		panic("too strong")
	}else {
		fmt.Println("input value is ",str)
	}
}