package s3

import (
	logger "api-gateway/pkg/logging"
	LoggerTypes "api-gateway/pkg/logging/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

var AWSConnection s3.S3

func InitS3Connection() {
	awsSession := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(os.Getenv("AWS_REGION")),
		Endpoint: aws.String(os.Getenv("AWS_ENDPOINT")),
	}))
	svc := s3.New(awsSession)
	AWSConnection = *svc
	logger.Log(LoggerTypes.INFO, "[AWS S3 | Connection] S3 connection initialized", nil)
}
