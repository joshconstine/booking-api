// objectStorage/client.go

package objectStorage

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"booking-api/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

type S3Client struct {
	Service *s3.S3
	Bucket  string
}

var Client *S3Client
var objectStorageError error

func CreateSession() {
	env, err := config.LoadConfig("../")
	if err != nil {
		fmt.Printf("error: %v", err)

		objectStorageError = err
		return
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-ord-1"),
		Endpoint:    aws.String("https://us-ord-1.linodeobjects.com"),
		Credentials: credentials.NewStaticCredentials(env.OBJECT_STORAGE_ACCESS_KEY, env.OBJECT_STORAGE_SECRET, ""),
	})
	if err != nil {
		fmt.Printf("error: %v", err)

		objectStorageError = err
		return

	}

	Client = &S3Client{
		Service: s3.New(sess),
		Bucket:  env.OBJECT_STORAGE_BUCKET,
	}

}

func (client *S3Client) UploadFile(ctx context.Context, key string, file http.File, filename string, fileExt string) (string, error) {
	// Generate a unique filename using a UUID
	newFilename := uuid.New().String() + fileExt

	// Construct new file path
	newFilePath := filepath.Join(key, newFilename)

	_, err := client.Service.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(client.Bucket),
		Key:         aws.String(newFilePath),
		Body:        file,
		ACL:         aws.String("public-read"),
		ContentType: aws.String(fileExt),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			return "", fmt.Errorf("upload canceled due to timeout: %w", err)
		}
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	return newFilePath, nil
}
