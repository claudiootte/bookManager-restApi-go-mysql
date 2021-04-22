# Book manager Rest API Go/MySQL
## API Restful que utiliza livros como exemplo

<p align="center">
  <img src="images/golang.png"/ alt="Golang">
</p>


A aplicação foi escrita totalmente em Go 🐹, visando utilizar o mínimo de dependências possíveis, tratar a maioria dos erros com os devidos cuidados e documentada com comentários de fácil entendimento

Pacotes utilizados

- Gorilla/Mux 🦍
- GORM


## Features

- Na aplicação é possível ver os livros já adicionados previamente
- Adicionar novos livros
- Ver livros específicos
- Excluir algum livros desejado
- Atualizar as informações


## Requisitos

```sh
Golang: https://golang.org/dl/
MySQL instalado e configurado com as configurações desejadas (editar no arquivo credentials.go)
API Client de sua preferência (O que aparece nas imagens se chama Postman)

Após ter instalado o go
Gorilla mux: go get -u github.com/gorilla/mux
GORM: go get -u gorm.io/gorm
```


## Na prática

Iniciando a aplicação

![](images/1-starting.png)


Acessando os livros já inseridos

![](images/2-getbooks.png)


Adicionando novos livros

![](images/3-addbooks01.png)

![](images/3-addbooks02.png)

![](images/3-addbooks03check.png)

![](images/3-addbooks03check.png)


Atualizando o livro desejado

![](images/4-updatebooks01.png)

![](images/4-updatebooks02check.png)


Deletando o livro desejado

![](images/5-deletebooks01wrong.png)

![](images/5-deletebooks02.png)

![](images/5-deletebooks03check.png)



## Utilização

Basta utilizar o comando go build e aproveitar o aplicativo! 😊
