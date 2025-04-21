# 📌 Ponteiros em Go

Em Go, **ponteiros** são variáveis que armazenam o **endereço de memória** de outra variável. Usando ponteiros, podemos manipular diretamente o valor da variável a partir de seu endereço de memória, o que é especialmente útil para modificar grandes estruturas de dados ou objetos sem precisar copiá-los.

---

## 🧠 Conceitos Básicos de Ponteiros

### Declaração de Ponteiro

Para declarar um ponteiro, usamos o símbolo `*` na frente do tipo da variável. O símbolo `&` é utilizado para obter o endereço de memória de uma variável.

```go
var a int = 10
var ponteiro *int = &a // ponteiro é um ponteiro para um inteiro
println(ponteiro)       // Exibe o endereço de memória de 'a'
```

### Dereferência de Ponteiro

A **dereferência** é o processo de acessar o valor armazenado no endereço de memória de um ponteiro. Isso é feito usando o `*` na frente do ponteiro.

```go
println(*ponteiro) // Exibe o valor armazenado no endereço de 'ponteiro', que é 10
```

### Alterando o Valor da Variável Através do Ponteiro

Ao modificar o valor da variável através do ponteiro, o valor original da variável também é alterado, já que o ponteiro aponta para o endereço de memória dessa variável.

```go
variavel := 10
abc(&variavel)  // Passando o endereço da variável
println(variavel) // Exibe 200, pois o valor foi alterado através do ponteiro

// Função que altera o valor apontado pelo ponteiro
func abc(a *int) {
	*a = 200
}
```

### Uso de Ponteiros em Métodos

Em Go, podemos usar **ponteiros em métodos** para alterar os valores das estruturas de dados diretamente. Isso é feito ao usar um receiver **por ponteiro** no método.

#### Passando o Receiver por Valor

Quando passamos a estrutura **por valor** em um método, estamos criando uma cópia da estrutura. Alterações feitas dentro do método **não afetam** a estrutura original.

```go
type Aluno struct {
	nome string
}

func (a Aluno) falou() {
	a.nome = "Lucas" // A alteração não afeta o valor original de a.nome
	println("Oi, meu nome é", a.nome)
}

func main() {
	Aluno1 := Aluno{
		nome: "Caio",
	}

	Aluno1.falou()       // Exibe "Oi, meu nome é Lucas"
	println(Aluno1.nome) // Exibe "Caio", pois a alteração foi feita na cópia
}
```

#### Passando o Receiver por Referência (Ponteiro)

Quando usamos um **ponteiro** como receiver do método (`*Aluno`), estamos passando o endereço de memória da estrutura, permitindo que as alterações feitas dentro do método afetem diretamente a estrutura original.

```go
func (a *Aluno) falou() {
	a.nome = "Lucas" // Agora, a alteração afeta o valor original de a.nome
	println("Oi, meu nome é", a.nome)
}

func main() {
	Aluno1 := Aluno{
		nome: "Caio",
	}

	Aluno1.falou()       // Exibe "Oi, meu nome é Lucas"
	println(Aluno1.nome) // Exibe "Lucas", pois a alteração foi feita diretamente na estrutura original
}
```

---

## 🧑‍🏫 Explicando o Processo

### Passando por Valor

Quando passamos uma estrutura por valor para um método, o Go cria uma cópia dessa estrutura, e qualquer alteração feita dentro do método não altera o valor original.

### Passando por Referência (Ponteiro)

Quando passamos uma estrutura por referência (usando ponteiros), o Go passa o **endereço de memória** da estrutura para o método. Isso significa que qualquer alteração feita no método afetará diretamente a estrutura original.
