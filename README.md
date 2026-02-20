# Agenda Telefônica

Esta é uma aplicação de Agenda Telefônica Full-Stack, com backend em Go e frontend em React (utilizando Vite), orquestrada inteiramente via Docker Compose.

## Pré-requisitos

Para executar a aplicação, você só precisa ter as seguintes ferramentas instaladas em sua máquina:

- **Docker:** Incluindo Docker Compose.

## Configuração do Ambiente

1.  **Clone o Repositório:**
    ```bash
    git clone [URL_DO_SEU_REPOSITÓRIO] # Substitua pela URL real do repositório
    cd app-agenda-telefonica
    ```

2.  **Configure as Variáveis de Ambiente:**
    No diretório `backend/`, renomeie o arquivo `.env.example` para `.env` e configure as variáveis necessárias.
    ```bash
    cd backend
    cp .env.example .env
    ```
    Edite o arquivo `backend/.env` com suas configurações, especialmente para o banco de dados. Um exemplo básico seria:
    ```
    # backend/.env
    DB_USER=postgres
    DB_PASS=postgres
    DB_HOST=db # Nome do serviço do banco de dados no docker-compose.yml
    DB_PORT=5432
    DB_NAME=agenda
    API_PORT=8080
    LOG_PATH=logs/exclusao.log
    ```
    *   Certifique-se de que `DB_HOST` corresponde ao nome do serviço do banco de dados definido no seu arquivo `docker-compose.yml`.
    *   O `API_PORT` no backend deve corresponder à porta exposta pelo container do backend no `docker-compose.yml`.

## Como Executar a Aplicação

Na raiz do projeto (onde se encontra o arquivo `docker-compose.yml`), execute o seguinte comando:

```bash
docker-compose up --build -d
```

*   `--build`: Garante que as imagens Docker sejam construídas a partir dos `Dockerfile`s.
*   `-d`: Executa os containers em segundo plano.

Este comando iniciará todos os serviços necessários (backend, frontend e banco de dados) configurados no `docker-compose.yml`.

## Acessando a Aplicação

*   **Frontend:** Acesse a aplicação no seu navegador em `http://localhost:[PORTA_DO_FRONTEND]` (a porta exata dependerá da sua configuração no `docker-compose.yml`).
*   **Backend API:** A API do backend estará disponível em `http://localhost:[PORTA_DO_BACKEND]` (a porta exata dependerá da sua configuração no `docker-compose.yml`, geralmente 8080).

## Testando a API do Backend

O script `test_api.sh` na raiz do projeto pode ser usado para testar a API do backend.

1.  Torne o script executável:
    ```bash
    chmod +x test_api.sh
    ```
2.  Execute o teste:
    ```bash
    ./test_api.sh
    ```
    *Nota: Certifique-se de que o script `test_api.sh` esteja configurado para se conectar à porta correta do backend exposta pelo Docker Compose.*

## Estrutura de Pastas

- `backend/`: Contém o código fonte do backend em Go, `Dockerfile`, `go.mod`, `migrations/`, e arquivos de configuração (`.env.example`).
- `frontend/`: Contém o código fonte do frontend (React/Vite).
- `docker-compose.yml`: Define a orquestração de todos os serviços (backend, frontend, banco de dados).
- `README.md`: Este arquivo.
- `test_api.sh`: Script para testar a API do backend.

## Documentação da API do Backend

| Método    | Rota               | Descrição                                     | Payload (Exemplo)                                                                                            |
| ---------- | ------------------ | ----------------------------------------------- | ------------------------------------------------------------------------------------------------------------ |
| `POST`   | `/contatos`      | Cria um novo contato e seus telefones.          | `{"id": 1, "nome": "Nome", "idade": 30, "telefones": [{"id_contato": 1, "id": 1, "numero": "912345678"}]}` |
| `GET`    | `/contatos`      | Lista todos os contatos e seus telefones.       | N/A                                                                                                          |
| `GET`    | `/contatos/{id}` | Busca um único contato pelo seu ID.            | N/A                                                                                                          |
| `PUT`    | `/contatos/{id}` | Atualiza um contato existente e seus telefones. | `{"nome": "Novo Nome", "idade": 31, "telefones": [{"id_contato": 1, "id": 1, "numero": "987654321"}]}`     |
| `DELETE` | `/contatos/{id}` | Deleta um contato pelo seu ID.                  | N/A                                                                                                          |
