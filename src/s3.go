package main

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// See https://docs.aws.amazon.com/sdk-for-go/api/service/s3/ for documentation.
func uploadToS3(reader io.ReadSeeker, s3bucket string, s3key string) error {
	svc := GeAwsS3Service("us-west-1")
	params := s3.PutObjectInput{
		Bucket: aws.String(s3bucket),
		Key:    aws.String(s3key),
		Body:   reader,
	}
	_, err := svc.PutObject(&params)
	return err
}
