package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// These are passed through by GitHub based on the inputs
	// Prefix is always INPUT_ and then it's the input name in uppercase
	bucket := os.Getenv("INPUT_BUCKET")
	region := os.Getenv("INPUT_REGION")
	sourceFiles := os.Getenv("INPUT_SOURCE")

	if bucket == "" || region == "" || sourceFiles == "" {
		log.Fatalf("expected INPUT_BUCKET, INPUT_REGION & INPUT_SOURCE to be set as env vars")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	svc := s3.NewFromConfig(cfg)

	// Upload each file to S3 after trimming the root directory from the prefix
	err = filepath.WalkDir(sourceFiles, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			childFile := strings.TrimPrefix(path, fmt.Sprintf("%s/", sourceFiles))

			_, err = svc.PutObject(context.Background(), &s3.PutObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(childFile),
			})

			if err != nil {
				log.Fatalf("problem uploading file %s: %v", childFile, err)
			}
		}

		return nil
	})
	if err != nil {
		log.Fatalf("problem uploading files %s", err)
	}
}
