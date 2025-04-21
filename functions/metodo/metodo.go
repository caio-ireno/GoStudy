package metodo

type Aluno struct {
	Nota int
}

func (a Aluno) CalcularNota() (bool, int) {
	if a.Nota >= 6 {
		return true, a.Nota // Aprovado
	} else {
		return false, a.Nota // Reprovado
	}
}
