#!/bin/bash

# Script para testar a API da Agenda Telefônica
# Certifique-se de que a aplicação Go esteja rodando antes de executar.

BASE_URL="http://localhost:8080"
LOG_FILE="backend/logs/exclusao.log"
TEST_COUNT=0
FAIL_COUNT=0

# --- Funções de Ajuda ---
print_test_name() {
    ((TEST_COUNT++))
    echo "---"
    echo "▶️  Test $TEST_COUNT: $1"
}

assert_status() {
    local expected=$1
    local actual=$2
    local test_name=$3

    if [ "$actual" -eq "$expected" ]; then
        echo "✅  PASS: Status esperado ($expected) recebido."
    else
        echo "❌  FAIL: Status esperado ($expected), mas recebeu ($actual)."
        echo "   (Teste: $test_name)"
        ((FAIL_COUNT++))
    fi
}

# Limpa o log de exclusão para um novo teste
rm -f $LOG_FILE

# --- Início dos Testes ---

# 1. Criar um Contato com sucesso
print_test_name "Criar Contato (Sucesso)"
response_code=$(curl -s -o /dev/null -w "%{http_code}" -X POST "$BASE_URL/contatos" \
-H "Content-Type: application/json" \
-d '{
    "id": 101,
    "nome": "Fulano de Tal",
    "idade": 25,
    "telefones": [{"id_contato": 101, "id": 1, "numero": "99999-0001"}]
}')
assert_status 201 "$response_code" "Criar Contato (Sucesso)"

# 2. Forçar Erro de Requisição Inválida (JSON mal formatado)
print_test_name "Criar Contato (Erro 400 - Requisição Inválida)"
response_code=$(curl -s -o /dev/null -w "%{http_code}" -X POST "$BASE_URL/contatos" \
-H "Content-Type: application/json" \
-d '{
    "id": 102,
    "nome": "Ciclano",
    "idade": "idade_invalida"
}')
assert_status 400 "$response_code" "Criar Contato (Erro 400 - Requisição Inválida)"

# 3. Listar Contatos
print_test_name "Listar Todos os Contatos"
response_code=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/contatos")
assert_status 200 "$response_code" "Listar Todos os Contatos"

# 4. Buscar Contato por ID (Sucesso)
print_test_name "Buscar Contato por ID (Sucesso)"
response_code=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/contatos/101")
assert_status 200 "$response_code" "Buscar Contato por ID (Sucesso)"

# 5. Buscar Contato por ID (Não Encontrado)
print_test_name "Buscar Contato por ID (Erro 404 - Não Encontrado)"
response_code=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/contatos/9999")
assert_status 404 "$response_code" "Buscar Contato por ID (Erro 404 - Não Encontrado)"

# 6. Atualizar Contato (Sucesso)
print_test_name "Atualizar Contato (Sucesso)"
response_code=$(curl -s -o /dev/null -w "%{http_code}" -X PUT "$BASE_URL/contatos/101" \
-H "Content-Type: application/json" \
-d '{
    "nome": "Fulano de Tal ATUALIZADO",
    "idade": 26,
    "telefones": [{"id_contato": 101, "id": 1, "numero": "99999-1111"}]
}')
assert_status 200 "$response_code" "Atualizar Contato (Sucesso)"

# 7. Deletar Contato (Sucesso)
print_test_name "Deletar Contato (Sucesso)"
response_code=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$BASE_URL/contatos/101")
assert_status 204 "$response_code" "Deletar Contato (Sucesso)"

# 8. Verificar se o Contato foi Deletado
print_test_name "Verificar se o Contato foi Deletado (Erro 404)"
response_code=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/contatos/101")
assert_status 404 "$response_code" "Verificar se o Contato foi Deletado (Erro 404)"

# 9. Tentar Deletar um Contato Inexistente
print_test_name "Tentar Deletar Contato Inexistente (Erro 500)"
# A API retorna 500 porque o repositório avisa que o ID não foi encontrado para deleção
response_code=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE "$BASE_URL/contatos/9999")
assert_status 500 "$response_code" "Tentar Deletar Contato Inexistente (Erro 500)"

# 10. Verificar se o Log de Exclusão foi Criado e Contém o ID
print_test_name "Verificar Arquivo de Log de Exclusão"
if [ -f "$LOG_FILE" ]; then
    echo "✅  PASS: Arquivo de log '$LOG_FILE' encontrado."
    if grep -q "Contato ID 101 excluído" "$LOG_FILE"; then
        echo "✅  PASS: Log para o contato ID 101 encontrado no arquivo."
    else
        echo "❌  FAIL: Log para o contato ID 101 NÃO encontrado no arquivo."
        ((FAIL_COUNT++))
    fi
else
    echo "❌  FAIL: Arquivo de log '$LOG_FILE' NÃO encontrado."
    ((FAIL_COUNT++))
fi

# --- Resumo Final ---
echo "---"
if [ "$FAIL_COUNT" -eq 0 ]; then
    echo "🎉 SUCESSO! Todos os $TEST_COUNT testes passaram."
else
    echo "🚨 ATENÇÃO: $FAIL_COUNT de $TEST_COUNT testes falharam."
fi
echo "---"
exit $FAIL_COUNT
