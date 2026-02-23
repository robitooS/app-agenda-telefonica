# Agenda Telefônica Full-Stack

Esta é uma aplicação de Agenda Telefônica completa, desenvolvida com backend em Go e frontend em React (TypeScript), orquestrada eficientemente com Docker Compose. Projetada com foco em robustez, inclui tratamento de erros aprimorado e validação de schema para garantir a integridade dos dados.

## Funcionalidades Principais

* **Gerenciamento de Contatos (CRUD):** Crie, visualize, atualize e delete contatos.
* **Gerenciamento de Telefones:** Adicione múltiplos telefones a cada contato.
* **Pesquisa Dinâmica:** Busque contatos por nome e/ou número de telefone.
* **Tratamento de Erros Profissional:** Respostas de API padronizadas e seguras, evitando vazamento de detalhes internos.
* **Integridade Referencial:** Deleção em cascata para telefones, garantida pelo banco de dados.

## Tecnologias Utilizadas

### Backend (Go)

* **Linguagem:** Go
* **Framework Web:** Gin Gonic
* **Banco de Dados:** PostgreSQL (via `database/sql` e `pq` driver)
* **Gerenciamento de Dependências:** Go Modules
* **Tratamento de Erros:** Pacote `errors` padrão e customizado (`internal/errors`)
* **Log:** Pacote `log` padrão

### Frontend (React com TypeScript)

* **Framework:** React
* **Linguagem:** TypeScript
* **Build Tool:** Vite
* **HTTP Client:** Axios
* **Ícones:** Lucide React
* **Estilização:** CSS Vanilla

### Orquestração e Ambiente

* **Containerização:** Docker
* **Orquestração:** Docker Compose

## Pré-requisitos

Para executar a aplicação, você só precisa ter as seguintes ferramentas instaladas em sua máquina:

* **Docker:** Incluindo **Docker Compose V2**.
* **Git:** Para clonar o repositório.

## Configuração e Instalação

1. **Clone o Repositório:**

   ```bash
   git clone https://github.com/robitooS/app-agenda-telefonica.git
   cd app-agenda-telefonica
   ```
2. **Configure as Variáveis de Ambiente do Backend:**
   No diretório `backend/`, renomeie o arquivo `.env.example` para `.env` e configure as variáveis.

   ```bash
   cd backend
   cp .env.example .env
   ```

## Como Executar a Aplicação

Na raiz do projeto (onde se encontra o arquivo `docker-compose.yml`), execute o seguinte comando:

```bash
docker compose up --build
```

* `--build`: Garante que as imagens Docker sejam construídas a partir dos `Dockerfiles`.
* Este comando iniciará e orquestrará todos os serviços necessários (backend, frontend e banco de dados).

### Detalhes Importantes:

* **Migrações do Banco de Dados:** O backend aplica automaticamente as migrações do banco de dados na inicialização, garantindo que o schema esteja sempre atualizado.
* **Integridade Referencial:** A deleção de um `Contato` resultará na deleção automática de todos os seus `Telefones` associados, graças à cláusula `ON DELETE CASCADE` configurada no schema do banco de dados.

## Acessando a Aplicação

* **Frontend:** Acesse a aplicação no seu navegador em `http://localhost:5173`.
* **Backend API:** A API do backend estará disponível em `http://localhost:8080`.

## Formato de Erro Padronizado

Em caso de erro, a API retornará uma resposta JSON no seguinte formato:

```json
{
  "code": "CODIGO_DO_ERRO",
  "message": "Mensagem amigavel para o usuario",
  "details": ["Detalhe adicional 1", "Detalhe adicional 2"]
}
```

Exemplos de códigos de erro: `NAO_ENCONTRADO`, `ENTRADA_INVALIDA`, `JA_EXISTE`, `ERRO_INTERNO_SERVE`.

## Estrutura de Pastas

* `backend/`: Contém o código fonte do backend em Go, `Dockerfile`, `go.mod`, `migrations/`, e arquivos de configuração (`.env.example`).
* `frontend/`: Contém o código fonte do frontend (React/Vite), `Dockerfile`.
* `docker-compose.yml`: Define a orquestração de todos os serviços (backend, frontend, banco de dados).
* `README.md`: Este arquivo.
