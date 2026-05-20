package main

import (
	"fmt"
	"soma-simples/pkg/soma"
)

func main()  {
	resultado := soma.Soma(10, 5)
	fmt.Printf("10 + 5 = %d\n", resultado)
}