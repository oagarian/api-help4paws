# API_Help4Paws

API para o aplicaitvo [Help4Paws](https://github.com/oagarian/Help4Paws)

## Como rodar

### 1. Clone o repositório e entre diretório do projeto

~~~bash
$ git clone https://github.com/oagarian/API_app.git
$ cd API_app/
~~~

### 2. Configure as variáveis de variáveis de ambiente

Crie uma cópia do arquivo `.env.example` com o nome de `.env`.

E preencha utilziando as variáveis informadas no arquivo mencionado.

### 3.  Instale as dependências e execute a api

~~~bash
$ go mod tidy
$ cd src
$ go run ./
~~~

<br>

![go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![postgres](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
