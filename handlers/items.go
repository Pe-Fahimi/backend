package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ketabdoozak/backend/middlewares"
	"github.com/ketabdoozak/backend/models"
	"github.com/ketabdoozak/backend/pkg/db"
	"github.com/ketabdoozak/backend/pkg/responses"
	"net/http"
)

func ListItems() gin.HandlerFunc {
	type Response struct {
		Results []models.Item `json:"results"`
		Total   int           `json:"total"`
	}
	return func(ctx *gin.Context) {
		var res []models.Item
		err := db.DB().Where("status = ?", models.ItemStatusPublished).Find(&res).Error
		if err != nil {
			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error(), Message: "error on finding items"})
			return
		}

		ctx.JSON(http.StatusOK, Response{Results: res, Total: len(res)})
	}
}

func CreateItem() gin.HandlerFunc {
	type Request struct {
		Title      string `json:"title" binding:"required"`
		Content    string `json:"content"`
		LocationID int64  `json:"location_id" binding:"required"`
		CategoryID int64  `json:"category_id" binding:"required"`
	}

	return func(ctx *gin.Context) {
		session := middlewares.MustGetSession(ctx)

		// get request
		var req Request
		err := ctx.Bind(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error(), Message: "invalid request data"})
			return
		}

		item := models.Item{
			Title:      req.Title,
			Content:    req.Content,
			AuthorID:   session.User.ID,
			LocationID: req.LocationID,
			CategoryID: req.CategoryID,
		}

		err = db.DB().Create(&item).Error
		if err != nil {
			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error(), Message: "error on creating item"})
			return
		}

		ctx.JSON(http.StatusCreated, item)
	}
}

func ReadItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get id
		id := ctx.Param("id")

		var item models.Item

		err := db.DB().Where("id = ?", id).First(&item).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				ctx.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "item not found"})
				return
			}

			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error(), Message: "error on finding item"})
			return
		}

		ctx.JSON(http.StatusOK, item)
	}
}

func UpdateItem() gin.HandlerFunc {
	type Request struct {
		Title      string `json:"title" binding:"required"`
		Content    string `json:"content"`
		LocationID int64  `json:"location_id" binding:"required"`
		CategoryID int64  `json:"category_id" binding:"required"`
	}

	return func(ctx *gin.Context) {
		// get id
		id := ctx.Param("id")

		// get request
		var req Request
		err := ctx.Bind(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error(), Message: "invalid request data"})
			return
		}

		var item models.Item

		err = db.DB().Where("id = ?", id).First(&item).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				ctx.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "item not found"})
				return
			}

			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error(), Message: "error on finding item"})
			return
		}

		if item.Title != req.Title {
			item.Status = models.ItemStatusPending
		}
		item.Title = req.Title
		if item.Content != req.Content {
			item.Status = models.ItemStatusPending
		}
		item.Content = req.Content
		item.LocationID = req.LocationID
		item.CategoryID = req.CategoryID

		err = db.DB().Save(&item).Error
		if err != nil {
			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error(), Message: "error on updating item"})
			return
		}

		ctx.JSON(http.StatusOK, item)
	}
}

func RemoveItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get id
		id := ctx.Param("id")

		var item models.Item

		err := db.DB().Where("id = ?", id).First(&item).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				ctx.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "item not found"})
				return
			}

			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error(), Message: "error on finding item"})
			return
		}

		err = db.DB().Delete(&item).Error
		if err != nil {
			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error(), Message: "error on deleting item"})
			return
		}

		ctx.JSON(http.StatusNoContent, responses.EmptyResponse{})
	}
}
