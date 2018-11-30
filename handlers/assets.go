package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ketabdoozak/backend/pkg/minio"
	"github.com/ketabdoozak/backend/pkg/responses"
	theMinio "github.com/minio/minio-go"
	"github.com/rs/xid"
	"io/ioutil"
	"net/http"
)

func UploadAsset() gin.HandlerFunc {
	type Response struct {
		ObjectName string `json:"object_name"`
	}
	return func(ctx *gin.Context) {
		objectName := xid.New().String()

		fh, err := ctx.FormFile("upload")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, responses.Error{Error: err.Error(), Message: "error on reading file from request"})
			return
		}

		f, err := fh.Open()
		if err != nil {
			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.Error{Error: err.Error(), Message: "error on opening file"})
			return
		}

		_, err = minio.Storage().PutObjectWithContext(ctx, minio.BucketName(), objectName, f, fh.Size, theMinio.PutObjectOptions{})
		if err != nil {
			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.Error{Error: err.Error(), Message: "error on putting object to storage"})
			return
		}

		ctx.JSON(http.StatusCreated, Response{ObjectName: objectName})
	}
}

func GetAsset() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		objectName := ctx.Param("object_name")
		if objectName == "" {
			ctx.JSON(http.StatusBadRequest, responses.Info{Message: "no object name entered in url"})
			return
		}
		obj, err := minio.Storage().GetObjectWithContext(ctx, minio.BucketName(), objectName, theMinio.GetObjectOptions{})
		if err != nil {
			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.Error{Error: err.Error(), Message: "error on getting object"})
			return
		}
		defer func() { _ = obj.Close() }()

		objectInfo, err := obj.Stat()
		if err != nil {
			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.Error{Error: err.Error(), Message: "error on getting object info"})
			return
		}

		data, err := ioutil.ReadAll(obj)
		if err != nil {
			// TODO: Submit error
			ctx.JSON(http.StatusInternalServerError, responses.Error{Error: err.Error(), Message: "error on reading object content"})
			return
		}

		ctx.Data(http.StatusOK, objectInfo.ContentType, data)
	}
}
