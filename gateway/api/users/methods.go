package users

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	dbErr "github.com/muazwzxv/go-backend-masterclass/db/errors"
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
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	})
	if err != nil {
		errCode := dbErr.ErrorCode(err)
		if errCode == dbErr.UniqueViolation {
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ErrorResponse(err))
		}

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
		if err == usersModule.ErrNotFound {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, utils.ToResponseBody(convertToResponseUser(user)))
}

func (h *Handler) LoginUser(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Errorf("h.LoginUser: %v", err)
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(utils.BadRequest))
		return
	}

	res, err := h.m.LoginUser(ctx, &usersModule.LoginUserRequest{
		UserName: req.Username,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(usersModule.ErrNotFound, err) {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(utils.NotFound))
			return
		}

		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.InternalServer))
	}

  ctx.JSON(http.StatusOK, LoginUserResponse{LoginResponse: res})
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	// TODO:
}

func (h *Handler) UpdatePassword(ctx *gin.Context) {
	// TODO:
}
