# EuLabs | Case Dev. Backend

## Pré-requisitos
Para executar os próximos passos é necessário ter o docker e docker-compose instalado.

## Começando

```
git clone https://github.com/adrianosiq/eulabs-challenge-api && cd eulabs-challenge-api
```

Para rodar os testes

* `go test ./...`
* `go test ./... -v`

Executando o servidor local

* `docker compose up app`

Agora visite [`localhost:8080`](http://localhost:8080) no seu navegador.

## Tecnologias
* Linguagem: Golang
* Framework: Echo Framework
* Database: MySQL 8
* Arquitetura da API: Rest
* CI: GitActions
