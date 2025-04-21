# 📘 Guia de Estudos em Go

Bem-vindo ao repositório de estudos em Go! Aqui você encontra os principais conceitos da linguagem organizados de forma simples e prática. 👨‍💻

## 🧭 Navegação

- [📌 Tipos e Variáveis](#-tipos-e-variáveis-em-go)
- [🧩 Funções e Uso de Pacotes](#-funções-e-uso-de-pacotes-em-go)
- [🚪 Exportação com Letra Maiúscula](#-exportação-de-funções-e-variáveis-em-go)

---

# 📌 Tipos e Variáveis em Go

Go é uma linguagem **estaticamente tipada com suporte à inferência de tipo**.

> ✅ O tipo da variável é verificado em tempo de compilação e **não muda em tempo de execução**.

---

## 🧠 Exemplo de tipo estático

```go
var idade int = 25
idade = "vinte e cinco" // ❌ Erro! Não pode mudar de int para string
```

---

## 🔍 Inferência de Tipo

Go consegue deduzir o tipo com base no valor atribuído, usando `:=` dentro de funções:

```go
nome := "Caio"
idade := 27
altura := 1.80
maiorIdade := true
```

> ⚠️ **Atenção:** A inferência só funciona com `:=` **dentro** de funções. Fora delas, você deve usar `var`.<br>
> ⚠️ **Atenção:** O uso `:=` é **apenas** para **criar** e **atribuir**. Para reatribuição use **apenas** `=`.

```go
nome := "João" // ❌ ERRO fora da função main ou outra função
```

Correto:

```go
var nome = "João"
```

---

## 🚫 Tipos não são dinâmicos

```go
nome := "João"
nome = 123 // ❌ Erro! Não posso trocar string por int
```

### Para comparação, veja como seria em Python:

```python
nome = "João"
nome = 123  # ✅ Ok em Python, mas proibido em Go
```

---

## 💡 Dica

Se quiser declarar uma variável **sem atribuir um valor inicial**, você **deve informar o tipo**:

```go
var sobrenome string
var idade int
```

---

## 🔢 Tipos de Dados Básicos

| Tipo                 | Descrição                          |
| -------------------- | ---------------------------------- |
| `string`             | Texto                              |
| `int`                | Número inteiro (plataforma padrão) |
| `int32`, `int64`     | Inteiros com tamanhos específicos  |
| `float32`, `float64` | Números com ponto flutuante        |
| `bool`               | Booleano: `true` ou `false`        |

### Exemplo:

```go
nome := "Maria"
idade := 30
altura := 1.65
ativo := true
```

---

## 🧩 Funções e Uso de Pacotes em Go

Em Go, funções são definidas com a palavra-chave `func` e devem pertencer a um **pacote** (`package`). Todo arquivo `.go` deve começar com a declaração do pacote ao qual ele pertence.

### 🔧 Exemplo de função simples

Arquivo: `variaveis/matematica/operations.go`

```go
package matematica

// Soma realiza a soma de dois inteiros
func Soma(a, b int) int {
  return a + b
}
```

---

### 📦 Estrutura de pacotes

A estrutura de diretórios pode representar pacotes. Exemplo:

```
gostudy/
│
├── go.mod
│
├── variaveis/
│   ├── main.go
│   └── matematica/
│       └── operations.go
```

---

### 📥 Importando um pacote interno

Para importar um pacote dentro do mesmo módulo, você usa o nome do módulo (definido no `go.mod`) seguido do caminho relativo:

> Se seu `go.mod` tem `module gostudy`, então:

```go
import (
  "fmt"
  "gostudy/variaveis/matematica"
)
```

---

### ▶️ Exemplo de uso no `main.go`

Arquivo: `variaveis/main.go`

```go
package main

import (
  "fmt"
  "gostudy/variaveis/matematica"
)

func main() {
  a := 10
  b := 5
  soma := matematica.Soma(a, b)
  fmt.Println("Soma:", soma)
}
```

> ⚠️ Importante: sempre execute seus programas a partir da raiz do módulo (onde está o `go.mod`).  
> Exemplo:

```bash
go run variaveis/main.go
```

---

## 🚪 Exportação de Funções e Variáveis em Go

Go tem uma regra simples, mas poderosa, para controle de **visibilidade**:

> ✅ **Começou com letra maiúscula? É exportado (público).**  
> ❌ **Começou com letra minúscula? É interno ao pacote (privado).**

---

### 📦 Variáveis em Pacotes

Arquivo: `matematica/operations.go`

```go
package matematica

var InternalVar string = "visível"
var internalVar string = "interno"
```

No `main.go`:

```go
fmt.Println(matematica.InternalVar)   // ✅ funciona
fmt.Println(matematica.internalVar)   // ❌ erro! não exportado
```

---

### ✅ Boas práticas

- Use letra maiúscula apenas para itens que precisam ser acessados de fora.
- Para manter a privacidade, prefira letra minúscula.
- Se quiser expor uma variável interna, crie uma função pública:

```go
func GetInternalVar() string {
  return internalVar
}
```

---
