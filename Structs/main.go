package main

import (
	"encoding/json"
	"fmt"
)

type Client struct {
	Name  string `json:"nome"`
	Age   int 	`json:"idade"`
	Email string `json:"email"`
}

func (c Client) Exibe(){
	fmt.Println("Nome: ", c.Name)
}

type ClientInternacional struct {
	Client
	Pais string `json:"pais"`
	Passport string `json:"passport"`
}

func main() {

	cliente := Client{
		Name:  "Caio",
		Age:   25,
		Email: "lucar@g.com",
	}

	fmt.Println(cliente)

	cliente2 := Client{"Nathalia", 24, "n@g.com"}
	fmt.Println(cliente2)
	fmt.Printf("Nome: %s\n", cliente2.Name)

	cliente3:= ClientInternacional{
		Client: Client{
			Name:  "Lucas",
			Age:   30,
			Email: "l@g.com",
		},
		Pais:     "EUA",
		Passport: "123456789",
	}

	fmt.Println(cliente3)
	fmt.Printf("Nome: %s, Idade: %d, Pais: %s, Passport: %s\n", cliente3.Name, cliente3.Age, cliente3.Pais, cliente3.Passport)

	cliente3.Exibe()


	clienteJson, err := json.Marshal(cliente3)
	if err != nil{
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	fmt.Println("Cliente em JSON:", string(clienteJson))

	jsonCLiente4 := `{"nome":"Lucas","idade":30,"email":"l@g.com","pais":"EUA","passport":"123456789"}`
	cliente4 := ClientInternacional{}

	json.Unmarshal([]byte(jsonCLiente4), cliente4)
	fmt.Println("Cliente 4:", cliente4) //se usar apenas cliente 4, ele vai alterar apenas nesse escopo, mas não no cliente4 original

	json.Unmarshal([]byte(jsonCLiente4), &cliente4)
	fmt.Println("Cliente 4:", cliente4) //se usar o &cliente4, ele vai alterar o cliente4 original, pois está passando o endereço de memória dele

}