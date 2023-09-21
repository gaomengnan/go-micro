package api

import (
	"net/http"
	"time"

	db "github.com/gaomengnan/go-micro/db/sqlc"
	"github.com/gaomengnan/go-micro/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type createUserReponse struct {
	Username          string    `db:"username" json:"username"`
	FullName          string    `db:"full_name" json:"full_name"`
	Email             string    `db:"email" json:"email"`
	PasswordChangedAt time.Time `db:"password_changed_at" json:"password_changed_at"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var params = new(createUserRequest)

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(params.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username:       params.Username,
		HashedPassword: hashedPassword,
		FullName:       params.FullName,
		Email:          params.Email,
	}
	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		if pdErr, ok := err.(*pq.Error); ok {
			switch pdErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createUserReponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}
