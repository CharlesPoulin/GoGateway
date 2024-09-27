package infra

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Client wraps the AWS S3 client
type S3Client struct {
	client *s3.Client
	bucket string
}

// NewS3Client creates and returns a new S3 client
func NewS3Client(bucket string) (*S3Client, error) {
	// Load the default configuration from environment variables or ~/.aws
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	// Create S3 client
	client := s3.NewFromConfig(cfg)

	return &S3Client{
		client: client,
		bucket: bucket,
	}, nil
}
