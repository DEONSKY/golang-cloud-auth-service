package s3service

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/forfam/authentication-service/src/utils/logger"
)

func getConfig() *aws.Config {
	accessKey := os.Getenv("AWS_S3_ACCESS_KEY_ID")
	if len(accessKey) == 0 {
		logger.GlobalLogger.Fatal(`Missing env variable "AWS_S3_ACCESS_KEY_ID"`)
	}

	secretKey := os.Getenv("AWS_S3_SECRET_KEY")
	if len(secretKey) == 0 {
		logger.GlobalLogger.Fatal(`Missing env variable "AWS_S3_SECRET_KEY"`)
	}

	region := os.Getenv("AWS_S3_REGION")
	if len(region) == 0 {
		logger.GlobalLogger.Fatal(`Missing env variable "AWS_S3_REGION"`)
	}

	server := os.Getenv("AWS_S3_SERVER")
	if len(server) == 0 {
		logger.GlobalLogger.Fatal(`Missing env variable "AWS_S3_SERVER"`)
	}

	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")
	_, err := creds.Get()
	if err != nil {
		logger.GlobalLogger.Fatal("Bad AWS S3 credentials")
	}

	// For debug AWS S3 requests
	// return aws.
	// 	NewConfig().
	// 	WithRegion(region).
	// 	WithCredentials(creds).
	// 	WithEndpoint(server).
	// 	WithLogLevel(aws.LogDebugWithHTTPBody)

	return aws.
		NewConfig().
		WithRegion(region).
		WithCredentials(creds).
		WithEndpoint(server)
}

var s3Client *s3.S3

func init() {
	logger.GlobalLogger.Info("S3 instance creation started")
	sess := session.New()
	s3Client = s3.New(sess, getConfig())

	_, err := s3Client.ListBuckets(nil)

	if err != nil {
		logger.GlobalLogger.Fatal("Something went wrong with S3 connection: " + err.Error())
	}

	logger.GlobalLogger.Info("S3 instance creation finished successfuly!")
}
