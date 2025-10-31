package repository

import (
	"context"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioRepository interface {
	UploadFile(ctx context.Context, bucketName, objectName string, fileBytes []byte, contentType string) (string, error)
	GetFileURL(ctx context.Context, bucketName, objectName string) (string, error)
	DeleteFile(ctx context.Context, bucketName, objectName string) error
}

type MinioRepositoryImpl struct {
	client *minio.Client
}

func MinioRepositoryInit(endpoint, accessKeyID, secretAccessKey string, useSSL bool) (*MinioRepositoryImpl, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return &MinioRepositoryImpl{
		client: minioClient,
	}, nil
}

func (m *MinioRepositoryImpl) UploadFile(ctx context.Context, bucketName, objectName string, file multipart.File, fileSize int64, contentType string) (string, error) {
	_, err := m.client.PutObject(
		ctx,
		bucketName,
		objectName,
		file,
		fileSize,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return "", err
	}
	url := m.client.EndpointURL().String() + "/" + bucketName + "/" + objectName
	return url, nil
}

func (m *MinioRepositoryImpl) GetFileURL(ctx context.Context, bucketName, objectName string) (string, error) {
	reqParams := make(url.Values)

	presignedURL, err := m.client.PresignedGetObject(ctx, bucketName, objectName, time.Duration(1000)*time.Second, reqParams)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}
