package url

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	var body CreateURLParams

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, body)
}

func (h *Handler) GetByShortCode(c *gin.Context) {
	// TODO: get short code from URL param, call h.service.GetByShortCode(), return response
}

func (h *Handler) Delete(c *gin.Context) {
	// TODO: get short code from URL param, call h.service.Delete(), return response
}
