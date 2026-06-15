# Roadmap de Evolução da GoAPI

Este documento detalha os próximos passos para a evolução da API, ordenados pela prioridade técnica e arquitetural de implementação.

## 1. Configurações e Variáveis de Ambiente (Fundação)
**Por que fazer primeiro:** Antes de conectar banco de dados ou criar tokens JWT, precisamos de um lugar seguro para guardar strings de conexão e chaves secretas, tirando o que está "hardcoded" no código.
- [x] Instalar pacote para ler arquivos `.env` (Ex: `github.com/joho/godotenv` ou `spf13/viper`).
- [x] Criar arquivo `.env` (adicionado ao `.gitignore`) e um `.env.example`.
- [x] Mudar a porta do servidor `:8080` no `main.go` para ler de uma variável de ambiente `PORT`.

## 2. Refatoração de Arquitetura - Injeção de Dependências
**Por que fazer agora:** Atualmente, a API chama `tools.NewDatabase()` a cada requisição HTTP. Se colocarmos um banco de dados real agora, isso vai abrir múltiplas conexões simultâneas até travar o banco.
- [x] Inicializar o banco de dados **apenas uma vez** no `main.go`.
- [x] Refatorar os Handlers (como `GetAccountBalance`) para receberem a instância do banco (usando uma *Struct* para o handler).
- [x] Refatorar o `AuthorizationMiddleware` para também receber a instância do banco de dados injetada.

## 3. Banco de Dados Real (PostgreSQL)
**Por que fazer agora:** Com as configurações seguras e a arquitetura arrumada, podemos substituir o `MockDatabase` por um banco de dados de verdade sem riscos estruturais.
- [x] Escolher um driver/ferramenta (ex: `lib/pq` puro, GORM ou sqlc).
- [x] Instalar e configurar o banco PostgreSQL (e go) com docker.
- [x] Criar a conexão com o banco no código usando as variáveis do `.env` (ex: `DATABASE_URL`).
- [x] Implementar migrations (ex: `golang-migrate/migrate`) para criar as tabelas `users` e `accounts`.
- [x] Atualizar a interface `DatabaseInterface` para interagir com o PostgreSQL no lugar do Mock.

## 4. Autenticação e Segurança (JWT)
**Por que fazer agora:** Agora que temos o banco real para validar usuários, podemos abandonar os tokens fixos.
- [x] Alterar o banco de dados para armazenar senhas como *hashes* (usando `golang.org/x/crypto/bcrypt`).
- [x] Criar o endpoint de Login (`POST /login`) que verifica credenciais e gera o JWT.
- [x] Definir a chave secreta do JWT (`JWT_SECRET`) no arquivo `.env`.
- [x] Atualizar o `AuthorizationMiddleware` para extrair, validar a assinatura e ler o `account_id` de dentro do token JWT enviado no cabeçalho `Authorization: Bearer <token>`.

## 5. Novas Features de Negócio (Transações Financeiras)
**Por que fazer agora:** A base está sólida. É hora de brincar com as regras de negócio complexas.
- [ ] Implementar endpoint de Depósito (`POST /account/deposit`).
- [ ] Implementar endpoint de Saque (`POST /account/withdraw`).
- [ ] Implementar endpoint de Transferência (`POST /account/transfer`).
- [ ] Garantir atomicidade nas transferências usando transações no banco (`tx.Begin()`, `tx.Commit()`, `tx.Rollback()`) para evitar perda de dinheiro em caso de falha.
- [ ] **Histórico de Transações (Extrato):** Criar tabela `transactions` para registrar todas as transferências/depósitos/saques e expor um endpoint `GET /account/statement`.

## 6. Infraestrutura e Qualidade (Testes, Docker e Swagger)
**Por que fazer no fim (ou paralelamente):** Refinar o código para garantir que ele esteja pronto para produção.
- [ ] **Testes:** Escrever testes unitários e de integração (usando `testing`, `httptest` e o pacote `testify` que já está no go.mod).
- [x] **Docker:** Criar um `Dockerfile` (multistage build) para a API e um `docker-compose.yml` para subir o PostgreSQL e a API com um único comando.
- [ ] **Documentação:** Configurar o Swagger (com `swaggo/swag`) para documentar e facilitar o teste dos endpoints gerados.

## 7. Refatoração de Banco e Cadastro (Identificação por E-mail)
**Por que fazer:** Tornar o sistema de cadastro e acesso mais robusto e próximo de um serviço real usando e-mail como chave única de login.
- [ ] **Nova Modelagem DB:** Criar tabelas `users` (id, nome, email, cpf, password_hash) e relacionar com a tabela `accounts` (com numero_conta gerado automaticamente e saldo).
- [ ] **Endpoint de Cadastro (`POST /register`):** Receber nome, email, CPF e senha, gerando automaticamente um número de conta exclusivo (ex: `10023-4`).
- [ ] **Login por E-mail:** Atualizar `POST /login` para autenticar usando `email` + `password` em vez de `account_id`.
- [ ] **Tela de Cadastro no Front-end:** Criar formulário de registro completo e elegante no tema escuro.

## 8. Features de Banco Avançadas (Realismo)
**Por que fazer:** Deixar o app com "cara de banco digital".
- [ ] **Agência e Conta:** Separar as contas por agências simuladas (Ex: Agência 0001) para fins de transferência.
- [ ] **Limite de Crédito/Cheque Especial:** Permitir que contas fiquem com saldo negativo até um limite aprovado no cadastro.
- [ ] **Extrato Avançado:** Permitir filtrar o extrato por tipo (entrada/saída) e intervalo de datas.
- [ ] **Área de Favoritos/Contatos:** Salvar contas para as quais o usuário transfere com frequência.