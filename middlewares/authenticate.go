package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ketabdoozak/backend/models"
	"github.com/ketabdoozak/backend/pkg/db"
	"github.com/ketabdoozak/backend/pkg/responses"
)

const SessionKey = "Session"

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userToken := ctx.GetHeader("User-Token")
		if userToken == "" {
			ctx.JSON(http.StatusUnauthorized, responses.Error{Message: "unauthorized"})
			ctx.Abort()
			return
		}

		var session models.Session
		err := db.DB().Where(models.Session{Token: userToken}).Preload("User").First(&session).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				ctx.JSON(http.StatusUnauthorized, responses.Error{Error: err.Error(), Message: "invalid user token"})
				ctx.Abort()
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, responses.Error{Error: err.Error(), Message: "error in checking session"})
				ctx.Abort()
				return
			}
		}

		if session.ExpiresAt.Before(time.Now()) {
			ctx.JSON(http.StatusUnauthorized, responses.Error{Error: err.Error(), Message: "session expired"})
			ctx.Abort()
			return
		}

		if session.DeletedAt != nil {
			ctx.JSON(http.StatusUnauthorized, responses.Error{Error: err.Error(), Message: "invalid user token"})
			ctx.Abort()
			return
		}

		ctx.Set(SessionKey, session)
		ctx.Next()
	}
}

func MustGetSession(ctx *gin.Context) models.Session {
	return ctx.MustGet(SessionKey).(models.Session)
}
