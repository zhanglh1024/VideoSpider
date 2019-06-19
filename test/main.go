package main

import "fmt"

func main() {
	defer func() {
		if err := recover();err!=nil{
			fmt.Println("some thing wrong")
		}
		var wrong string
		for {
			fmt.Println("please input the wrong info:")
			fmt.Scan(&wrong)
			if wrong == "exit"{
				panic("wrong")
			}
			fmt.Println(wrong)
		}
	}()
	var name, age, work string
	for {
		fmt.Scan(&name,&age, &work)
		if work == "exit"{
			panic("wrong")
		}
		fmt.Println(name)
	}
}