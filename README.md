# Go MVC

Simples projeto de estudo de Go.

## Estrutura

```
src
├── api
│   ├── controller
│   │   └── users_controller.go
│   ├── middleware
│   │   └── method_not_allowed.go
│   ├── request
│   │   ├── user_request.go
│   │   └── validate.go
│   ├── response
│   │   └── user_response.go
│   └── routes
│       ├── routes.go
│       └── users.go
├── errors
│   └── rest_err.go
└── model
```

# Tecnologias

- [MongoDB](https://www.mongodb.com/)
- [Golang 1.22+](https://golang.org/)
- [Gin Gonic](https://gin-gonic.com/docs/)
- [Validator](https://github.com/go-playground/validator)

# Features

- [ ] CRUD de usuários
  - [ ] Listar usuários
  - [ ] Buscar usuário por ID
  - [ ] Criar usuário
  - [ ] Atualizar usuário
  - [ ] Deletar usuário
- [ ] Validação de dados
- [ ] Testes unitários
- [ ] Testes de integração
- [ ] Autenticação JWT
- [ ] Autorização
- [ ] SWAGGER
- [ ] Documentação
