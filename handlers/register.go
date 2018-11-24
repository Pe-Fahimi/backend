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
	"regexp"
	"time"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func Register() gin.HandlerFunc {
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
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error(), Message: "invalid request data"})
			return
		}

		// check email is valid
		if !emailRegexp.MatchString(req.Email) {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid email format"})
			return
		}

		// check user exists
		var user models.User
		err = db.DB().Where(models.User{Email: req.Email}).First(&user).Error
		if err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				// TODO: Submit error
				ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error(), Message: "error on checking user"})
				return
			}
		} else {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "email already exists"})
			return
		}

		// user doesn't exist
		// create user
		user = models.User{}
		user.Email = req.Email
		user.PasswordHash, err = password.HashPassword(req.Password)
		if err != nil {
			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "error on creating password hash", Error: err.Error()})
			return
		}

		user.RegisteredAt = time.Now()
		err = db.DB().Create(&user).Error
		if err != nil {
			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "error on creating user", Error: err.Error()})
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
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error(), Message: "error on creating session"})
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
