package users

import (
	"github.com/gin-gonic/gin"
	"github.com/muazwzxv/go-backend-masterclass/modules/users"
	"go.uber.org/zap"
)

type Handler struct {
	m   users.IUsers
	log *zap.SugaredLogger
}

func New(
  module users.IUsers,
  log *zap.SugaredLogger,
) *Handler {
  return &Handler{
    m: module,
    log: log,
  }
}

func (h *Handler) Routes(route *gin.RouterGroup) {
  v1 := route.Group("/v1")
  {
    v1.GET("/user/:id", h.GetUser)
    v1.POST("/user", h.CreateUser)
  }
}
