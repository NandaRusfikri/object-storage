package main

import (
	"context"
	"log"
	"object-storage/helper"
)

func main() {
	ctx := context.Background()

	accessKey := "o9USFt70xrzMcBmxbDvp"
	SecretKey := "frYw8eoOzqAyAoJonpjg8wSGdZutDFvOYP7ZBLXl"
	bucket := "product"
	host := "localhost:9000"

	// Initialize the MinIO client
	minioClient, err := helper.NewMinioClient(host, accessKey, SecretKey, bucket, false)
	if err != nil {
		log.Fatalf("Error initializing Minio client: %v", err)
	}

	objectName := "img.jpg"
	filePath := "file/img.jpg"
	contentType := "image/jpg"
	//Upload file
	err = minioClient.UploadFile(ctx, objectName, filePath, contentType)
	if err != nil {
		log.Fatalf("Error uploading file: %v", err)
	}

	minioClient.GetObjectPresign(ctx, objectName)

	//// Download file
	//err = minioClient.DownloadFile(ctx, "example.txt", "/file/example.txt")
	//if err != nil {
	//	log.Fatalf("Error downloading file: %v", err)
	//}
	//
	//// Delete file
	//err = minioClient.DeleteFile(ctx, "example.txt")
	//if err != nil {
	//	log.Fatalf("Error deleting file: %v", err)
	//}
}
