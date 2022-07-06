# BACKEND implementada em GO

Implementação de backend monolitico, contendo microserviços incorporados, utilizando o GRPC com Clean Archetecture e TDD.

## Caracteristas

- [X] Testes  (TDD)
- [X] Design Patterns
- [X] Clean Architecture
- [X] Suporte para desligamento seguro
- [X] gRPC API
- [X] Log de produção
- [X] Cadastro de Conta do Usuário
- [X] Login
- [X] Autenticação
- [X] Conexão SSL/TLS

## Build Server

go build -o bin/server -buildmode pie ./cmd/monolithic

## Build Client (test)

go build -o bin/client -buildmode pie ./cmd/client

## Run server

./bin/server

## Add Account

./bin/client devaccount
