package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/robitooS/backend/internal/entity"
	"github.com/robitooS/backend/internal/service"
	"net/http"
	"strconv"
)

type ContatoHandler struct {
	service service.ContatoService
}

func NewContatoHandler(s service.ContatoService) *ContatoHandler {
	return &ContatoHandler{service: s}
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

	if err := h.service.Create(&contato); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, contato)
}

func (h *ContatoHandler) GetContatos(c *gin.Context) {
	contatos, err := h.service.FindAll()
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

	contato, err := h.service.FindByID(id)
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

	if err := h.service.Update(&contato); err != nil {
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

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
