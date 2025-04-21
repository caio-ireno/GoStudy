package main

import (
	"gostudy/functions/metodo"
)
func main() {
	res := func() {
		println("Hello, World!")
	}
	res() // Hello, World!

	res2 := func(x ...int) func() int {
		res := 0
		for _, v := range x {
			res += v
		}
		return func() int {
			return res
		}
	}

	res3 := res2(1, 2, 3, 4, 5)
	println(res2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)()) // 55
	println(res3())                                // 15

	resultado := soma(10, 20)
	println(resultado) // 30

	resultado2 := somaTudo(10, 20, 30)
	println(resultado2) // 60

	aluno1 := metodo.Aluno{Nota: 10}
	print(aluno1.CalcularNota()) // true 10
}

func soma(a int, b int) (result int) {
	result = a + b
	return
}

func somaTudo(x ...int) int {
	resultado := 0
	for _, v := range x {
		resultado += v
	}
	return resultado
}
