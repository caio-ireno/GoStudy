package main

func main() {
	a := 10

	println(&a)
	println(a) // 10

	var ponteiro *int = &a // ponteiro é um ponteiro para um inteiro
	println(ponteiro)      // endereço de a
	println(*ponteiro)     // 10

	*ponteiro = 20     // altera o valor de a através do ponteiro
	println(a)         // 20
	println(ponteiro)  // endereço de a
	println(*ponteiro) // 20

	b := &a
	println(b)  // endereço de a
	println(*b) // 20
	println(a)  // 20

	variavel := 10
	abc(&variavel)    // passando o endereço da variável
	println(variavel) // 200

	Aluno1 := Aluno{
		nome: "Caio",
	}

	Aluno1.falou()
	println(Aluno1.nome) // Oi, meu nome é Lucas
}

// nao retorna nada, mas altera o valor passado pelo parametro
func abc(a *int) {
	*a = 200
}

type Aluno struct {
	nome string
}

// o uso do * altera o comportamento do método, fazendo com que ele altere o valor do campo nome
// se não usar o * o valor do campo nome não será alterado
func (a *Aluno) falou() {
	a.nome = "Lucas"
	println("Oi, meu nome é", a.nome)
}