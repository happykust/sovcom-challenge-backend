package s3

import (
	logger "api-gateway/pkg/logging"
	LoggerTypes "api-gateway/pkg/logging/types"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"mime/multipart"
	"os"
)

func UploadFile(file *multipart.FileHeader) string {

	bucket := os.Getenv("AWS_BUCKET")
	key := file.Filename
	ctx := context.Background()

	fileBody, err := file.Open()
	if err != nil {
		logger.Log(LoggerTypes.ERROR, "[AWS S3 | Upload] Error opening file", err)
		return ""
	}

	_, err = AWSConnection.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   fileBody,
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("successfully uploaded file to %s/%s\n", bucket, key)
	return fmt.Sprintf("https://%s.%s/%s", bucket, os.Getenv("AWS_ENDPOINT"), key)
}
