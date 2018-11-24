package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ketabdoozak/backend/models"
	"github.com/ketabdoozak/backend/pkg/db"
	"github.com/ketabdoozak/backend/pkg/responses"
	"net/http"
)

func ListLocations() gin.HandlerFunc {
	type Response struct {
		Results []models.Location `json:"results"`
		Total   int               `json:"total"`
	}
	return func(ctx *gin.Context) {
		var res []models.Location
		err := db.DB().Find(&res).Error
		if err != nil {
			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err, Message: "error on finding locations"})
			return
		}

		ctx.JSON(http.StatusOK, Response{Results: res, Total: len(res)})
	}
}
