package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ketabdoozak/backend/middlewares"
	"github.com/ketabdoozak/backend/pkg/db"
	"github.com/ketabdoozak/backend/pkg/responses"
	"net/http"
)

func Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := middlewares.MustGetSession(ctx)
		err := db.DB().Delete(&session).Error
		if err != nil {
			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.Error{Error: err.Error(), Message: "error on deleting session"})
			return
		}

		ctx.JSON(http.StatusOK, responses.Info{Message: "logged out successfully"})
	}
}
