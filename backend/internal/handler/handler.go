package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/robitooS/backend/internal/entity"
	"github.com/robitooS/backend/internal/service"
	"github.com/robitooS/backend/internal/logger"
	"net/http"
	"strconv"
)

type ContatoHandler struct {
	service service.ContatoService
	delLogPath string
}

func NewContatoHandler(s service.ContatoService, delLogPath string) *ContatoHandler {
	return &ContatoHandler{service: s, delLogPath: delLogPath}
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	if err := h.service.Create(ctx, &contato); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contatos)
}

func (h *ContatoHandler) GetContatoByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	ctx := c.Request.Context()
	contato, err := h.service.FindByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contato)
}

func (h *ContatoHandler) UpdateContato(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	var contato entity.Contato
	if err := c.ShouldBindJSON(&contato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contato.ID = id // Garante que o ID da URL seja usado

	ctx := c.Request.Context()
	if err := h.service.Update(ctx, &contato); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contato)
}

func (h *ContatoHandler) DeleteContato(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	ctx := c.Request.Context()
	if err := h.service.Delete(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.LogDeletedContact(h.delLogPath, id)
	c.JSON(http.StatusNoContent, nil)
}
