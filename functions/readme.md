# 🧠 Funções e Métodos em Go

## 📌 Navegação

- [🧠 Funções em Go](#-funções-em-go)
- [🔁 Múltiplos Retornos](#-múltiplos-retornos)
- [🏷️ Retorno Nomeado](#-retorno-nomeado)
- [🎲 Funções Variádicas](#-funções-variádicas)
- [🧩 Funções Anônimas](#-funções-anônimas)
- [🔄 Funções que Retornam Funções](#-funções-que-retornam-outras-funções)
- [👨‍🏫 Métodos em Go](#-métodos-em-go)

## 🧠 Funções em Go

Funções seguem o padrão:

```go
func nome(parametros) retorno {
	// corpo
}
```

### Exemplo simples:

```go
func soma(a int, b int) int {
	return a + b
}
```

## 🔁 Múltiplos Retornos

Go permite retornar mais de um valor:

```go
func soma(a int, b int) (int, string) {
	return a + b, "somou!"
}
```

## 🏷️ Retorno Nomeado

Você pode nomear a variável de retorno:

```go
func soma(a int, b int) (result int) {
	result = a + b
	return
}
```

## 🎲 Funções Variádicas

Recebem múltiplos argumentos de mesmo tipo:

```go
func somaTudo(x ...int) int {
	resultado := 0
	for _, v := range x {
		resultado += v
	}
	return resultado
}
```

## 🧩 Funções Anônimas

Funções sem nome atribuídas a variáveis:

```go
res := func() {
	fmt.Println("Hello, World!")
}
res() // Executa
```

## 🔄 Funções que Retornam Outras Funções

```go
res2 := func(x ...int) func() int {
	res := 0
	for _, v := range x {
		res += v
	}
	return func() int {
		return res
	}
}

res3 := res2(1, 2, 3)
fmt.Println(res3()) // 6
```

## 👨‍🏫 Métodos em Go

Em Go, métodos são **funções associadas a tipos (structs)**.

### Exemplo:

```go
package metodo

type Aluno struct {
	Nota int
}

// Método associado à struct Aluno
func (a Aluno) CalcularNota() (bool, int) {
	if a.Nota >= 6 {
		return true, a.Nota // Aprovado
	}
	return false, a.Nota // Reprovado
}
```

### Uso no `main.go`:

```go
import (
	"fmt"
	"gostudy/functions/metodo"
)

func main() {
	aluno1 := metodo.Aluno{Nota: 10}
	aprovado, nota := aluno1.CalcularNota()
	fmt.Println(aprovado, nota) // true 10
}
```

> ✅ **Resumo**: Em Go, _método_ é uma função com um _receiver_ (ex: `func (a Aluno)`) e funciona como os métodos de classes em linguagens orientadas a objetos.
