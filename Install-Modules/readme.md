## ⚙️ Módulos e Instalação de Dependências em Go

Go usa um sistema de módulos para gerenciamento de pacotes e dependências.

### 🏁 Inicializando um Módulo Go

Antes de tudo, você deve inicializar um módulo Go na raiz do projeto:

```bash
go mod init gostudy
```

Isso cria um arquivo `go.mod`, que gerencia os pacotes usados no projeto.

---

### 📦 Instalando Dependências

Para instalar pacotes de terceiros, como o pacote de UUID do Google, use:

```bash
go get github.com/google/uuid@latest
```

> O Go vai atualizar os arquivos `go.mod` e `go.sum` automaticamente.

---

### ▶️ Como Rodar o Código

Para executar um dos arquivos do projeto, use:

```bash
go run ./Install-Modules/main.go
```

> 📌 Certifique-se de sempre rodar os arquivos a partir da **raiz do projeto**, onde está o `go.mod`.
