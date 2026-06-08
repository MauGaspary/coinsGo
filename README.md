# GoAPI - Sistema de Consulta de Saldo Bancário

Uma API financeira robusta construída em Go, focada em performance e manutenibilidade. Este projeto utiliza uma arquitetura limpa com injeção de dependências e acesso a banco de dados type-safe com SQLC.

## 🚀 Tecnologias Utilizadas

- **Linguagem:** Go (v1.23.2)
- **Roteador:** [Chi](https://github.com/go-chi/chi)
- **Banco de Dados:** PostgreSQL
- **Gerador de Código SQL:** [SQLC](https://sqlc.dev/)
- **Logging:** [Logrus](https://github.com/sirupsen/logrus)
- **Variáveis de Ambiente:** [Godotenv](https://github.com/joho/godotenv)
- **Containerização:** Docker & Docker Compose

## 🛠️ Arquitetura e Organização

O projeto segue uma estrutura organizada para facilitar a escalabilidade:

- `cmd/api/`: Ponto de entrada da aplicação.
- `api/`: Definições de tipos e handlers de erro globais.
- `internal/handlers/`: Lógica de processamento das requisições HTTP.
- `internal/database/`: Código gerado pelo SQLC e lógica de acesso ao banco.
- `internal/middleware/`: Middlewares de segurança e autorização.
- `sql/`: Scripts de schema e queries SQL.
- `static/`: Interface web simples (HTML/CSS/JS).

## 📋 Funcionalidades Atuais

- [x] Servidor de arquivos estáticos para interface web.
- [x] Middleware de autorização básico.
- [x] Consulta de saldo de conta bancária.
- [x] Integração com PostgreSQL via Docker.
- [x] Injeção de dependências para melhor testabilidade.

## 🚦 Como Executar

### Pré-requisitos
- Docker e Docker Compose instalados.
- Go 1.23+ (opcional, se usar apenas Docker).

### Passo a Passo

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/MauGaspary/goapi.git
   cd goapi
   ```

2. **Configure as variáveis de ambiente:**
   Crie um arquivo `.env` na raiz do projeto seguindo o modelo do `.env.example`.
   ```bash
   PORT=8080
   DATABASE_URL=postgres://user:pass@db:5432/goapi?sslmode=disable
   ```

3. **Inicie a aplicação com Docker Compose:**
   ```bash
   docker-compose up --build
   ```

4. **Acesse a API:**
   - Interface Web: `http://localhost:8080`
   - Endpoint de Saldo: `GET http://localhost:8080/account/balance?account_id=SEU_ID`

## 🛠️ Desenvolvimento

### Banco de Dados (SQLC)
O projeto utiliza o [SQLC](https://sqlc.dev/) para gerar código Go a partir de queries SQL. Se você alterar os arquivos em `sql/`, execute:
```bash
sqlc generate
```

### Docker
Para reconstruir os containers após alterações no código:
```bash
docker-compose up --build
```

## 🔌 Endpoints da API

### `GET /account/balance`
Retorna o saldo atual de uma conta.

**Parâmetros:**
- `account_id` (Query String): ID da conta a ser consultada.

**Headers:**
- `Authorization`: Token de autenticação.

**Resposta de Sucesso (200 OK):**
```json
{
  "balance": 1500.50,
  "code": 200
}
```

## 🗄️ Modelo de Dados

### Tabela `accounts`
| Coluna | Tipo | Descrição |
|--------|------|-----------|
| `account_id` | VARCHAR(255) | Identificador único (PK) |
| `balance` | NUMERIC(20,4) | Saldo da conta |

### Tabela `login_details`
| Coluna | Tipo | Descrição |
|--------|------|-----------|
| `account_id` | VARCHAR(255) | FK para accounts (PK) |
| `password_hash` | TEXT | Hash da senha para autenticação |

---
Desenvolvido por [MauGaspary](https://github.com/MauGaspary)
