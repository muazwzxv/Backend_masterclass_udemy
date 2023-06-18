package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muazwzxv/go-backend-masterclass/gateway/utils"
  usersModule "github.com/muazwzxv/go-backend-masterclass/modules/users"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
  var req CreateUserRequest
  if err := ctx.ShouldBindJSON(&req); err != nil {
    h.log.Errorf("h.CreateUser: %v", err)
    ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.BadRequest))
    return
  }

  user, err := h.m.CreateUser(ctx, &usersModule.CreateUser{
    FirstName: req.FirstName,
    LastName: req.LastName,
    Email: req.Email,
  })
  if err != nil {
    ctx.AbortWithStatus(http.StatusInternalServerError)
    return
  }

  ctx.JSON(http.StatusCreated, utils.ToResponseBody(user))
}

func (h *Handler) GetUser(ctx *gin.Context) {
  var req GetUserRequest
  if err := ctx.ShouldBindUri(&req); err != nil {
    h.log.Errorf("h.GetUser: %v", err)
    ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.BadRequest))
  }

  user, err := h.m.FindUser(ctx, req.ID)
  if err != nil {
    h.log.Errorf("h.GetUser: %v", err)
    if err == usersModule.NotFound {
      ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
      return
    }

    ctx.AbortWithStatus(http.StatusInternalServerError)
    return
  }

  ctx.JSON(http.StatusOK, utils.ToResponseBody(convertToResponseUser(user)))
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
  // TODO:
}

func (h *Handler) UpdatePassword(ctx *gin.Context) {
  // TODO:
}
