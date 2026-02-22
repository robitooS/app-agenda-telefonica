package handler

import (
	"errors" // Importacao do pacote errors padrao do Go
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robitooS/backend/internal/entity"
	"github.com/robitooS/backend/internal/logger"
	"github.com/robitooS/backend/internal/service"

	errorsCustom "github.com/robitooS/backend/internal/errors"
)

type ContatoHandler struct {
	service service.ContatoService
	delLogPath string
}

func NewContatoHandler(s service.ContatoService, delLogPath string) *ContatoHandler {
	return &ContatoHandler{service: s, delLogPath: delLogPath}
}

func (h *ContatoHandler) handleError(c *gin.Context, err error) {
	var apiError *errorsCustom.APIError
	// Tenta desembrulhar o erro para ver se é um dos nossos erros customizados
	if errors.Is(err, errorsCustom.ErrNotFound) {
		apiError = errorsCustom.NewAPIError("NAO_ENCONTRADO", "Recurso nao encontrado", err.Error())
		c.JSON(http.StatusNotFound, apiError)
		return
	}
	if errors.Is(err, errorsCustom.ErrInvalidInput) {
		apiError = errorsCustom.NewAPIError("ENTRADA_INVALIDA", "Dados de entrada invalidos", err.Error())
		c.JSON(http.StatusBadRequest, apiError)
		return
	}
	if errors.Is(err, errorsCustom.ErrAlreadyExists) {
		apiError = errorsCustom.NewAPIError("JA_EXISTE", "Recurso ja existe", err.Error())
		c.JSON(http.StatusConflict, apiError)
		return
	}

	// Para outros erros (incluindo os wrapped de DB), retornar um erro interno genérico
	// Isso evita vazar detalhes internos para o cliente da API
	apiError = errorsCustom.NewAPIError("ERRO_INTERNO_SERVE", "Ocorreu um erro interno no servidor", "Por favor, tente novamente mais tarde.")
	c.JSON(http.StatusInternalServerError, apiError)
	log.Printf("Erro interno: %v", err)
}

func (h *ContatoHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/contatos", h.CreateContato)
	router.GET("/contatos", h.GetContatos)
	router.GET("/contatos/:id", h.GetContatoByID)
	router.PUT("/contatos/:id", h.UpdateContato)
	router.DELETE("/contatos/:id", h.DeleteContato)
}

func (h *ContatoHandler) CreateContato(c *gin.Context) {
	var contato entity.Contato
	if err := c.ShouldBindJSON(&contato); err != nil {
		h.handleError(c, errorsCustom.WrapErrorf(err, "entrada invalida para criacao de contato"))
		return
	}

	ctx := c.Request.Context()
	if err := h.service.Create(ctx, &contato); err != nil {
		h.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, contato)
}

func (h *ContatoHandler) GetContatos(c *gin.Context) {
	nome := c.Query("nome")
	numero := c.Query("numero")

	ctx := c.Request.Context()
	contatos, err := h.service.FindWithFilters(ctx, nome, numero)
	if err != nil {
		h.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, contatos)
}

func (h *ContatoHandler) GetContatoByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.handleError(c, errorsCustom.WrapErrorf(err, "entrada invalida para ID do contato"))
		return
	}

	ctx := c.Request.Context()
	contato, err := h.service.FindByID(ctx, id)
	if err != nil {
		h.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, contato)
}

func (h *ContatoHandler) UpdateContato(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.handleError(c, errorsCustom.WrapErrorf(err, "entrada invalida para ID do contato"))
		return
	}

	var contato entity.Contato
	if err := c.ShouldBindJSON(&contato); err != nil {
		h.handleError(c, errorsCustom.WrapErrorf(err, "entrada invalida para atualizacao de contato"))
		return
	}
	contato.ID = id // Garante que o ID da URL seja usado

	ctx := c.Request.Context()
	if err := h.service.Update(ctx, &contato); err != nil {
		h.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, contato)
}

func (h *ContatoHandler) DeleteContato(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.handleError(c, errorsCustom.WrapErrorf(err, "entrada invalida para ID do contato"))
		return
	}

	ctx := c.Request.Context()
	if err := h.service.Delete(ctx, id); err != nil {
		h.handleError(c, err)
		return
	}
	logger.LogDeletedContact(h.delLogPath, id)
	c.JSON(http.StatusNoContent, nil)
}
