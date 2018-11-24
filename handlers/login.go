package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ketabdoozak/backend/models"
	"github.com/ketabdoozak/backend/pkg/db"
	"github.com/ketabdoozak/backend/pkg/password"
	"github.com/ketabdoozak/backend/pkg/responses"
	"github.com/rs/xid"
	"net/http"
	"time"
)

func Login() gin.HandlerFunc {
	type Request struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	type Response struct {
		SessionID int64     `json:"session_id"`
		UserID    int64     `json:"user_id"`
		Token     string    `json:"token"`
		CreatedAt time.Time `json:"created_at"`
		ExpiresAt time.Time `json:"expires_at"`
	}

	return func(ctx *gin.Context) {
		// get request
		var req Request
		err := ctx.Bind(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err, Message: "invalid request data"})
			return
		}

		// find user
		var user models.User
		err = db.DB().Where(models.User{Email: req.Email}).First(&user).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err, Message: "user does not exist"})
			} else {
				// TODO: Submit error
				ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err, Message: "error on finding user"})
				return
			}
		}

		// check password is valid
		if !password.CheckPasswordHash(req.Password, user.PasswordHash) {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err, Message: "invalid user password"})
			return
		}

		// create session
		session := models.Session{}
		session.UserID = user.ID
		session.Token = xid.New().String()
		clientIP := ctx.ClientIP()
		session.ClientIP = &clientIP
		userAgent := ctx.Request.UserAgent()
		session.UserAgent = &userAgent
		session.CreatedAt = time.Now()
		session.ExpiresAt = time.Now().AddDate(1, 0, 0) // Expires after 1 years
		err = db.DB().Create(&session).Error
		if err != nil {
			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err, Message: "error on creating session"})
			return
		}

		session.User = &user

		res := Response{
			SessionID: session.ID,
			UserID:    user.ID,
			Token:     session.Token,
			CreatedAt: session.CreatedAt,
			ExpiresAt: session.ExpiresAt,
		}

		ctx.JSON(http.StatusOK, res)
	}
}
