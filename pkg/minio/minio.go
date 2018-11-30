package minio

import (
	"github.com/minio/minio-go"
	"github.com/pkg/errors"
	"log"
	"os"
)

var storage *minio.Client

func Storage() *minio.Client {
	if storage == nil {
		var err error
		secure := os.Getenv("MINIO_SECURE") == "1"
		storage, err = minio.New(
			os.Getenv("MINIO_ENDPOINT"),
			os.Getenv("MINIO_ACCESS_KEY"),
			os.Getenv("MINIO_SECRET_KEY"),
			secure,
		)
		if err != nil {
			log.Panic(errors.Wrap(err, "error on getting new minio client"))
		}

		bucketLocation := os.Getenv("MINIO_BUCKET_LOCATION")

		var exists bool
		exists, err = storage.BucketExists(BucketName())

		if err != nil {
			log.Panic(errors.Wrapf(err, "error on checking existence for minio bucket '%s'", BucketName()))
		}

		if !exists {
			err = storage.MakeBucket(BucketName(), bucketLocation)
			if err != nil {
				log.Panic(errors.Wrapf(err, "error on creating minio bucket '%s' in location '%s'", BucketName(), bucketLocation))
			}
		}
	}

	return storage
}

func BucketName() string {
	return os.Getenv("MINIO_BUCKET_NAME")
}
