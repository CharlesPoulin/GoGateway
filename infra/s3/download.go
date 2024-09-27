package infra

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// DownloadFile downloads a file from the specified S3 bucket
func (s *S3Client) DownloadFile(key, destination string) error {
	// Get the file from the S3 bucket
	output, err := s.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
	})
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}
	defer output.Body.Close()

	// Create a local file to save the downloaded content
	file, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Copy the content from S3 to the local file
	_, err = io.Copy(file, output.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}
