package helper

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"net/url"
	"time"
)

type MinioClient struct {
	Client *minio.Client
	Bucket string
}

// Initialize Minio client
func NewMinioClient(endpoint, accessKeyID, secretAccessKey, bucketName string, useSSL bool) (*MinioClient, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize minio client: %v", err)
	}

	//retentionMode := minio.Compliance
	//valid := uint(10)
	//ValidityUnit := minio.Days
	//err = client.SetBucketObjectLockConfig(context.Background(), bucketName, &retentionMode, &valid, &ValidityUnit)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	return &MinioClient{
		Client: client,
		Bucket: bucketName,
	}, nil
}

// Create (Upload) file to bucket
func (m *MinioClient) UploadFile(ctx context.Context, objectName, filePath, contentType string) error {
	// Upload the file to MinIO bucket
	_, err := m.Client.FPutObject(ctx, m.Bucket, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return fmt.Errorf("failed to upload file: %v", err)
	}
	log.Printf("Successfully uploaded %s to bucket %s\n", objectName, m.Bucket)
	return nil
}

// Get (Download) file from bucket
func (m *MinioClient) DownloadFile(ctx context.Context, objectName, destFilePath string) error {
	err := m.Client.FGetObject(ctx, m.Bucket, objectName, destFilePath, minio.GetObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}

	log.Printf("Successfully downloaded %s from bucket %s\n", objectName, m.Bucket)
	return nil
}

// Delete file from bucket
func (m *MinioClient) DeleteFile(ctx context.Context, objectName string) error {
	err := m.Client.RemoveObject(ctx, m.Bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}
	log.Printf("Successfully deleted %s from bucket %s\n", objectName, m.Bucket)
	return nil
}

func (m *MinioClient) GetObjectPresign(ctx context.Context, objectName string) string {

	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "inline; filename=\"your-filename.txt\"")

	// Generates a presigned url which expires in a day.
	presignedURL, err := m.Client.PresignedGetObject(ctx, m.Bucket, objectName, time.Minute*60*60, reqParams)
	if err != nil {
		return fmt.Sprintf("failed to delete file: %v", err)
	}
	log.Println("Successfully generated presigned URL", presignedURL)

	return presignedURL.String()
}
