# ❌ Tratamento de Erros em Go

Esta aula demonstra como **lidar com erros de forma explícita** em Go, uma das características mais marcantes da linguagem. Go **não usa try/catch**, como outras linguagens — em vez disso, os erros são **valores retornados** pelas funções.

---

## 📦 Exemplo na prática

O código deste diretório mostra duas situações comuns:

1. Fazer uma requisição HTTP e verificar se houve erro.
2. Criar uma função personalizada (`soma`) que retorna erro se os valores forem negativos.

### 🔍 Código-fonte principal

```go
res, err := http.Get("http://google.com.br")
if err != nil {
	log.Fatal("Erro no request:", err.Error())
}
fmt.Println("Status Code:", res.Header)
```

```go
resultado, err := soma(10, 5)
if err != nil {
	log.Fatal("Erro na soma:", err.Error())
}
fmt.Println("Resultado da soma:", resultado)
```

---

## 💡 O que é `nil`?

Em Go, `nil` representa **"nada" ou "valor ausente"**. Quando uma função não retorna erro, ela retorna `nil` como segundo valor.

```go
return a + b, nil
```

Isso indica que **não houve erro**.

---

## 📘 Funções com erro personalizado

A função `soma` mostra como criar e retornar um erro personalizado usando `fmt.Errorf`:

```go
func soma(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("não pode ser negativo")
	}
	return a + b, nil
}
```

> ✨ Isso é útil para validar argumentos e proteger a lógica do programa.

---

## ✅ Boas Práticas com Erros

- Sempre verifique `err` com `if err != nil`
- Use `log.Fatal(err)` se quiser encerrar a execução
- Retorne `nil` quando **não houver erro**
- Crie mensagens de erro claras para ajudar no debug
- Use `fmt.Errorf` para incluir contexto no erro

## ⚠️ Ignorando erros com `_`

Em Go, é possível **ignorar valores retornados** por uma função usando o **underline** (`_`). Isso inclui o valor de erro:

```go
_, err := soma(10, -5)
```

Ou até mesmo:

```go
_, _ = soma(10, -5) // ❌ ignora completamente o erro
```

No entanto, você também pode ver código como este:

```go
_ = http.Get("http://google.com.br")
```

> Isso significa que o erro (ou outro valor de retorno) foi **intencionalmente ignorado**.

---

### 🚨 Por que isso é perigoso?

Ignorar erros pode ocultar problemas sérios no seu código. Alguns riscos incluem:

- ❌ Falhas silenciosas que você não percebe
- ❌ Dificuldade para debugar
- ❌ Programas funcionando parcialmente ou de forma inesperada
- ❌ Violação das boas práticas e filosofia do Go

---

### ✅ Quando (raramente) usar `_`

Use `_` **somente se tiver certeza absoluta** de que aquele erro não é relevante. Mesmo assim, é mais comum utilizar `_` para ignorar **valores não utilizados**, e **não erros**.

---

### 🧠 Dica

> Em Go, **ignorar erros é uma má prática, não um recurso recomendável.** Trate todos os erros, mesmo que seja para fazer um log simples.

---

## 📎 Execução

Para rodar este exemplo:

```bash
go run ./tratamento-erros/main.go
```

> Certifique-se de estar na raiz do projeto (onde está o `go.mod`).

---

## 📚 Links úteis

- [Documentação oficial sobre erros em Go](https://go.dev/blog/error-handling-and-go)
- [Pacote `errors`](https://pkg.go.dev/errors)
- [Pacote `log`](https://pkg.go.dev/log)
