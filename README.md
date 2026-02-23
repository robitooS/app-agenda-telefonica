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

* **Docker:** Incluindo **Docker Compose V2**. Verifique sua versão com `docker compose version`. Se encontrar problemas, consulte a seção "Solução de Problemas Comuns".
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

## Testando a API do Backend

O script `test_api.sh` na raiz do projeto pode ser usado para realizar testes básicos na API do backend.

1. Torne o script executável:

   ```bash
   chmod +x test_api.sh
   ```
2. Execute o teste:

   ```bash
   ./test_api.sh
   ```

   *Nota: Este script oferece uma verificação básica. Para um testing mais robusto, considere implementar testes unitários e de integração completos.*

## Documentação da API do Backend

A API do backend fornece os seguintes endpoints para gerenciamento de contatos:

### Formato de Erro Padronizado

Em caso de erro, a API retornará uma resposta JSON no seguinte formato:

```json
{
  "code": "CODIGO_DO_ERRO",
  "message": "Mensagem amigavel para o usuario",
  "details": ["Detalhe adicional 1", "Detalhe adicional 2"]
}
```

Exemplos de códigos de erro: `NAO_ENCONTRADO`, `ENTRADA_INVALIDA`, `JA_EXISTE`, `ERRO_INTERNO_SERVE`.

### Endpoints

| Método    | Rota               | Descrição                                                                                                                                                                             | Payload (Exemplo de Requisição)                                                            | Respostas Possíveis (Status HTTP)                                                                                                                                               |
| ---------- | ------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `POST`   | `/contatos`      | Cria um novo contato com seus telefones.                                                                                                                                                | `{"nome": "Maria Silva", "idade": 28, "telefones": [{"numero": "98765-4321"}]}`            | `201 Created` (Contato), `400 Bad Request` (Entrada inválida), `500 Internal Server Error` (Erro interno)                                                                 |
| `GET`    | `/contatos`      | Lista todos os contatos e seus telefones. Suporta filtros opcionais por `nome` e `numero` de telefone.                                                                              | N/A (query params:`/contatos?nome=Maria&numero=98765`)                                     | `200 OK` (Lista de Contatos), `500 Internal Server Error` (Erro interno)                                                                                                     |
| `GET`    | `/contatos/{id}` | Busca um único contato pelo seu ID.                                                                                                                                                    | N/A                                                                                          | `200 OK` (Contato), `404 Not Found` (Contato não encontrado), `400 Bad Request` (ID inválido), `500 Internal Server Error` (Erro interno)                              |
| `PUT`    | `/contatos/{id}` | Atualiza um contato existente e seus telefones. Todos os telefones existentes serão substituídos pelos fornecidos no payload.                                                         | `{"nome": "Maria Antunes", "idade": 29, "telefones": [{"id": 1, "numero": "99999-0000"}]}` | `200 OK` (Contato Atualizado), `404 Not Found` (Contato não encontrado), `400 Bad Request` (Entrada inválida/ID inválido), `500 Internal Server Error` (Erro interno) |
| `DELETE` | `/contatos/{id}` | Deleta um contato pelo seu ID. Esta operação também removerá automaticamente todos os telefones associados a ele, devido à configuração `ON DELETE CASCADE` no banco de dados. | N/A                                                                                          | `204 No Content` (Sucesso), `404 Not Found` (Contato não encontrado), `400 Bad Request` (ID inválido), `500 Internal Server Error` (Erro interno)                      |

## Solução de Problemas Comuns

### `ModuleNotFoundError: No module named 'distutils'` ao executar `docker-compose up`

Este erro ocorre geralmente quando o `docker-compose` (versão V1) é instalado via `pip` em um ambiente Python mais recente (Python 3.10+), que removeu o módulo `distutils`.

**Solução Recomendada:**

1. **Use o Docker Compose V2:** A versão V2 é um plugin do Docker CLI e é a forma moderna e recomendada de uso.
2. **Verifique sua instalação:**
   * Tente executar `docker compose version` (note o espaço entre `docker` e `compose`).
   * Se este comando funcionar, você já tem o V2 instalado.
3. **Se não tiver o V2 ou o comando acima não funcionar:**
   * **Desinstale qualquer `docker-compose` antigo:**
     ```bash
     sudo apt remove docker-compose # Para instalações via apt (Debian/Ubuntu)
     pip uninstall docker-compose # Para instalações via pip
     ```
   * **Instale o Docker Compose V2** (se já não veio com o Docker Desktop ou sua instalação do Docker Engine para Linux):
     ```bash
     # Crie a pasta de plugins se não existir
     sudo mkdir -p /usr/local/lib/docker/cli-plugins/
     # Baixe o binário (substitua a URL pela versão mais recente para sua arquitetura em https://github.com/docker/compose/releases)
     sudo curl -L https://github.com/docker/compose/releases/latest/download/docker-compose-linux-x86_64 -o /usr/local/lib/docker/cli-plugins/docker-compose
     # Torne-o executável
     sudo chmod +x /usr/local/lib/docker/cli-plugins/docker-compose
     ```
   * Após a instalação, use sempre `docker compose up --build -d` (com espaço).

## Estrutura de Pastas

* `backend/`: Contém o código fonte do backend em Go, `Dockerfile`, `go.mod`, `migrations/`, e arquivos de configuração (`.env.example`).
* `frontend/`: Contém o código fonte do frontend (React/Vite), `Dockerfile`.
* `docker-compose.yml`: Define a orquestração de todos os serviços (backend, frontend, banco de dados).
* `README.md`: Este arquivo.
* `test_api.sh`: Script para testes básicos da API do backend.
