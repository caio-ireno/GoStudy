# 🧱 Structs em Go

Em Go, uma `struct` é uma forma poderosa de agrupar dados relacionados. Apesar de lembrarem classes, structs **não fazem parte de uma abordagem orientada a objetos clássica** — Go **não é POO** (Go away 😄). No entanto, oferece recursos como composição e métodos associados que permitem organizar bem o código.

---

## 🧩 O que é uma Struct?

Uma `struct` em Go é uma estrutura de dados composta que permite agrupar variáveis de **diferentes tipos** dentro de uma mesma unidade.

### ❌ Dados soltos

```go
type ClientNome string
type ClientEmail string
type ClientPhone string
```

### ✅ Forma organizada com `struct`

```go
type Client struct {
	Name  string
	Age   int
	Email string
}

func main() {
	cliente := Client{
		Name:  "Caio",
		Age:   25,
		Email: "lucar@g.com",
	}
}
```

---

## 🔁 Composição (quase como herança)

Go não tem herança, mas permite **composição** de structs, o que é uma alternativa poderosa.

```go
type ClientInternacional struct {
	Client   // "herda" os campos da struct Client
	Pais     string
	Passport string
}
```

### Exemplo de uso:

```go
cliente3 := ClientInternacional{
	Client: Client{
		Name:  "Lucas",
		Age:   30,
		Email: "l@g.com",
	},
	Pais:     "EUA",
	Passport: "123456789",
}
```

---

## 🧠 Métodos em Structs

É possível **associar métodos a structs**. Mesmo se você criar o método em uma struct "pai", ele também funcionará para structs compostas.

```go
func (c Client) Exibe() {
	fmt.Println("Nome:", c.Name)
}
```

---

## 🧾 Trabalhando com JSON

Go permite **converter structs em JSON** com `json.Marshal` e **ler JSON para structs** com `json.Unmarshal`.

### 🔄 Marshal (Struct → JSON)

```go
clienteJson, err := json.Marshal(cliente3)
if err != nil {
	fmt.Println("Erro ao converter para JSON:", err)
	return
}
fmt.Println("Cliente em JSON:", string(clienteJson))
```

> O `Marshal` retorna um `[]byte`, por isso usamos `string(clienteJson)` para exibir.

---

## 🔖 Tags JSON

Go trabalha com **tags** para mapear os campos de uma struct para chaves personalizadas no JSON.

```go
type ClientInternacional struct {
	Client
	Pais     string `json:"pais"`
	Passport string `json:"passport"`
}
```

Assim, os nomes no JSON podem seguir convenções diferentes, como `camelCase` ou `snake_case`.

---

## 🔄 Unmarshal (JSON → Struct)

Transformando um JSON em uma struct:

```go
jsonCLiente4 := `{"nome":"Lucas","idade":30,"email":"l@g.com","pais":"EUA","passport":"123456789"}`
cliente4 := ClientInternacional{}

json.Unmarshal([]byte(jsonCLiente4), &cliente4)
fmt.Println("Cliente 4:", cliente4)
```

> ⚠️ Se você **não usar o `&` (endereço de memória)** no segundo argumento, a modificação só acontece localmente e não altera o valor original.
