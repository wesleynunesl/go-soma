package main

import "fmt"

func main()  {
	a := 10
	b := 6
	resultado := a + b
	if resultado != 15{
		fmt.Printf("O resultado de %d + %d é %d\n",a, b, resultado)
	}

	//fmt.Printf("%d + %d = %d\n", a, b, resultado)
}