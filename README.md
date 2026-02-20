# Agenda Telefônica

Este é um projeto de aplicação completa (Full-Stack) para uma Agenda Telefônica, demonstrando uma arquitetura moderna com Frontend, Backend e Banco de Dados. Todo o ambiente de desenvolvimento e produção é orquestrado via Docker Compose, garantindo que o usuário precise apenas do Docker instalado para rodar a aplicação inteira.

## Arquitetura

O projeto é construído com as seguintes tecnologias e componentes:

- **Frontend:** [Tecnologia do Frontend - Ex: React, Angular, Vue.js]
  * Interface de usuário interativa para gerenciar contatos e telefones.
  * (Assumindo porta 3000, 4200 ou 8080 dependendo da tecnologia)
- **Backend:** Go (Gin Framework)
  * API RESTful para manipulação de dados de contatos e telefones.
  * Implementa regras de negócio, persistência de dados e lógica de exclusão com log.
  * (Porta padrão: 8080)
- **Banco de Dados:** PostgreSQL
  * Armazenamento persistente para contatos e telefones.
  * Gerenciado por migrações para controle de versão do esquema.
- **Orquestração:** Docker Compose
  * Define e executa o ambiente multi-container com um único comando.

## Pré-requisitos

Para executar a aplicação completa, você precisará ter as seguintes ferramentas instaladas em sua máquina:

- **Git:** Para clonar o repositório.
- **Docker e Docker Compose:** Essenciais para construir, executar e gerenciar todos os serviços da aplicação (frontend, backend, banco de dados) de forma isolada e consistente.

## Como Executar a Aplicação Completa

Siga estes passos para ter o ambiente de desenvolvimento completamente funcional.

### Passo 1: Clone o Repositório

Abra seu terminal e clone este repositório para a sua máquina local:

```bash
git clone https://github.com/Higor-Paulino/app-agenda-telefonica.git
cd app-agenda-telefonica
```

### Passo 2: Arquivos de Configuração `.env`

Crie os arquivos de ambiente necessários para o backend e (futuramente) para o frontend.

* **Para o Backend:** O repositório já deve incluir `backend/.env` com os valores padrão. Se não existir, ou para sobrescrever, crie-o.

  Exemplo de `backend/.env` (assumindo que o `DB_HOST` será o nome do serviço no Docker Compose):

  ```
  DB_USER=postgres
  DB_PASS=postgres
  DB_HOST=localhost
  DB_PORT=5432
  DB_NAME=agenda
  API_PORT=8080
  LOG_PATH=logs/exclusao.log
  ```
* **Para o Frontend:** (Este arquivo será necessário quando o frontend for implementado)
  Crie um arquivo `frontend/.env` (ou similar) com as variáveis que o frontend precisará para se comunicar com o backend, por exemplo:

  ```
  REACT_APP_API_URL=http://localhost:8080
  ```

  *Atenção: O host `localhost` aqui é para o frontend rodando no seu navegador acessar o backend exposto na porta 8080 da sua máquina.*

### Passo 3: Inicie a Aplicação com Docker Compose

Na raiz do projeto, execute o seguinte comando para construir as imagens (se necessário) e iniciar todos os serviços da aplicação (frontend, backend e banco de dados) em segundo plano:

```bash
docker-compose up --build -d
```

* O `-d` roda os containers em segundo plano.
* O `--build` garante que as imagens mais recentes sejam construídas a partir dos Dockerfiles.

### Passo 4: Acesse a Aplicação

* **Frontend:** Uma vez que todos os serviços estejam rodando, acesse a aplicação frontend em seu navegador:
  `http://localhost:[PORTA_DO_FRONT]` (ex: `http://localhost:3000` se for React)
* **Backend API:** A API do backend estará disponível para requisições externas em:
  `http://localhost:8080`

## Testando a API do Backend

Você pode usar o script de teste para verificar as funcionalidades da API do backend.

Em um terminal na raiz do projeto, certifique-se de que o script tenha permissão de execução:

```bash
chmod +x test_api.sh
```

Depois, execute o teste:

```bash
./test_api.sh
```

## Estrutura de Pastas

- `backend/`: Contém o código fonte do backend em Go, `Dockerfile`, `go.mod`, etc.
- `frontend/`: Contém o código fonte do frontend (a ser implementado).
- `migrations/`: Contém os arquivos de migração SQL para o banco de dados.
- `docker-compose.yml`: Define a orquestração de todos os serviços.
- `test_api.sh`: Script para testar a API do backend.

## Documentação da API do Backend

| Método    | Rota               | Descrição                                     | Payload (Exemplo)                                                                                            |
| ---------- | ------------------ | ----------------------------------------------- | ------------------------------------------------------------------------------------------------------------ |
| `POST`   | `/contatos`      | Cria um novo contato e seus telefones.          | `{"id": 1, "nome": "Nome", "idade": 30, "telefones": [{"id_contato": 1, "id": 1, "numero": "912345678"}]}` |
| `GET`    | `/contatos`      | Lista todos os contatos e seus telefones.       | N/A                                                                                                          |
| `GET`    | `/contatos/{id}` | Busca um único contato pelo seu ID.            | N/A                                                                                                          |
| `PUT`    | `/contatos/{id}` | Atualiza um contato existente e seus telefones. | `{"nome": "Novo Nome", "idade": 31, "telefones": [{"id_contato": 1, "id": 1, "numero": "987654321"}]}`     |
| `DELETE` | `/contatos/{id}` | Deleta um contato pelo seu ID.                  | N/A                                                                                                          |
