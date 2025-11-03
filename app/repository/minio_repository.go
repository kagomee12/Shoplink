package repository

import (
	"context"
	"mime/multipart"
	"net/url"
	"shoplink/app/config"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioRepository interface {
	UploadFile(ctx context.Context, bucketName, objectName string, file multipart.File, fileSize int64, contentType string) (string, error)
	GetFileURL(ctx context.Context, bucketName, objectName string) (string, error)
	DeleteFile(ctx context.Context, bucketName, objectName string) error
}

type MinioRepositoryImpl struct {
	client *minio.Client
}

func MinioRepositoryInit(cfg *config.MinioConfig) *MinioRepositoryImpl {
	minioClient, err := minio.New(string(cfg.Endpoint), &minio.Options{
		Creds:  credentials.NewStaticV4(string(cfg.AccessKeyID), string(cfg.SecretAccessKey), ""),
		Secure: bool(cfg.UseSSL),
	})
	if err != nil {
		return nil
	}

	return &MinioRepositoryImpl{
		client: minioClient,
	}
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

func (m *MinioRepositoryImpl) DeleteFile(ctx context.Context, bucketName, objectName string) error {
	err := m.client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
