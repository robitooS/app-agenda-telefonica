package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

func LogDeletedContact(logFilePath string, contactID int64) {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("ERRO: Falha ao abrir o arquivo de log %s: %v", logFilePath, err)
		return // Não impede a operação principal se o log falhar
	}
	defer file.Close()

	logEntry := fmt.Sprintf("%s - Contato ID %d excluído.", time.Now().Format("2006-01-02 15:04:05"), contactID)
	if _, err := file.WriteString(logEntry); err != nil {
		log.Printf("ERRO: Falha ao escrever no arquivo de log %s: %v", logFilePath, err)
	}
}