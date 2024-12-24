package main

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"log"
	"mime"
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

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	svc := s3.NewFromConfig(cfg)

	// Upload each file to S3 after trimming the root directory from the prefix
	err = filepath.WalkDir(sourceFiles, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			s3Prefix := strings.TrimPrefix(path, fmt.Sprintf("%s/", sourceFiles))

			// Read the source file contents
			body, err := os.ReadFile(path)
			if err != nil {
				log.Fatalf("problem reading contents of file %s: %v", path, err)
			}
			reader := bytes.NewReader(body)

			// Calculate the MIME Content-type based on the extension
			ext := filepath.Ext(path)
			contentType := mime.TypeByExtension(ext)
			if contentType == "" {
				contentType = "application/octet-stream"
			}

			_, err = svc.PutObject(context.Background(), &s3.PutObjectInput{
				Bucket:      aws.String(bucket),
				Key:         aws.String(s3Prefix),
				Body:        reader,
				ContentType: aws.String(contentType),
			})
			if err != nil {
				log.Fatalf("problem uploading file %s: %v", path, err)
			}
		}

		return nil
	})
	if err != nil {
		log.Fatalf("problem uploading files %s", err)
	}

	// Add a GitHub Action output for the website URL
	url := fmt.Sprintf("url=http://%s.s3-website.%s.amazonaws.com", bucket, region)
	err = os.Setenv("GITHUB_OUTPUT", url)
	if err != nil {
		log.Fatalf("problem updating GITHUB_OUTPUT env var: %v", err)
	}
}
