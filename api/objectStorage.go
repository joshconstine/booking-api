package api

import (
	"booking-api/config"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

// S3 transaction
// turns type multipart.File into a byte array
func fileToBytes(file multipart.File) ([]byte, error) {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return fileBytes, nil
}

func UploadHandler(w http.ResponseWriter, r *http.Request, key string) (string, error) {
	var bucket string
	timeout := 20 * time.Second

	env, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	bucket = env.OBJECT_STORAGE_BUCKET

	ACCESS_KEY := env.OBJECT_STORAGE_ACCESS_KEY
	SECRET_KEY := env.OBJECT_STORAGE_SECRET

	config := &aws.Config{
		Region:      aws.String("us-ord-1"),
		Endpoint:    aws.String("https://us-ord-1.linodeobjects.com"), // Linode Object Storage endpoint
		Credentials: credentials.NewStaticCredentials(ACCESS_KEY, SECRET_KEY, ""),
	}

	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.

	sess := session.Must(session.NewSession(config))

	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.
	svc := s3.New(sess)

	// Create a context with a timeout that will abort the upload if it takes
	// more than the passed in timeout.
	ctx := context.Background()
	var cancelFn func()
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}
	// Ensure the context is canceled to prevent leaking.
	// See context package for more information, https://golang.org/pkg/context/
	if cancelFn != nil {
		defer cancelFn()
	}
	err = r.ParseMultipartForm(10 * 1024 * 1024) // 10 MB limit
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusInternalServerError)
		return "", err
	}

	file, header, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusInternalServerError)
		return "", err
	}
	defer file.Close()
	// Generate a unique filename using a UUID
	fileExt := filepath.Ext(header.Filename)
	newFilename := uuid.New().String() + fileExt

	// Create a new file in the "public/static" directory with the unique filename
	newFilePath := filepath.Join(key, newFilename)

	// Uploads the object to S3. The Context will interrupt the request if the
	// timeout expires.
	_, err = svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(newFilePath),
		Body:        file,
		ACL:         aws.String("public-read"), // Set ACL to public-read
		ContentType: aws.String(fileExt),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			// If the SDK can determine the request or retry delay was canceled
			// by a context the CanceledErrorCode error code will be returned.
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}
		os.Exit(1)
	}

	return newFilePath, nil

}

func DeleteHandler(w http.ResponseWriter, r *http.Request, key string) error {

	var bucket string
	timeout := 20 * time.Second

	env, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	bucket = env.OBJECT_STORAGE_BUCKET

	ACCESS_KEY := env.OBJECT_STORAGE_ACCESS_KEY
	SECRET_KEY := env.OBJECT_STORAGE_SECRET

	config := &aws.Config{
		Region:      aws.String("us-ord-1"),
		Endpoint:    aws.String("https://us-ord-1.linodeobjects.com"), // Linode Object Storage endpoint
		Credentials: credentials.NewStaticCredentials(ACCESS_KEY, SECRET_KEY, ""),
	}

	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.

	sess := session.Must(session.NewSession(config))

	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.
	svc := s3.New(sess)

	// Create a context with a timeout that will abort the upload if it takes
	// more than the passed in timeout.
	ctx := context.Background()
	var cancelFn func()
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}
	// Ensure the context is canceled to prevent leaking.
	// See context package for more information, https://golang.org/pkg/context/
	if cancelFn != nil {
		defer cancelFn()
	}

	// Delete the object from S3. The Context will interrupt the request if the
	// timeout expires.
	_, err = svc.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			// If the SDK can determine the request or retry delay was canceled
			// by a context the CanceledErrorCode error code will be returned.
			fmt.Fprintf(os.Stderr, "delete canceled due to timeout, %v\n", err)

		} else {
			fmt.Fprintf(os.Stderr, "failed to delete object, %v\n", err)
		}
		os.Exit(1)

	}

	return nil
}
