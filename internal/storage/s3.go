package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	cfgAws "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rvxt21/bucket-inventory/config"
)

type S3 struct {
	cfg    *config.Config
	client *s3.Client
	bucket string
}

func (s *S3) Start(ctx context.Context) error {
	cfg, err := cfgAws.LoadDefaultConfig(
		ctx,
		cfgAws.WithRegion(s.cfg.Region),
		cfgAws.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(s.cfg.AccessKey, s.cfg.SecretKey, "")))
	if err != nil {
		return err
	}

	s.client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		if s.cfg.Endpoint != "" {
			o.BaseEndpoint = aws.String(s.cfg.Endpoint)
			o.UsePathStyle = true
		}
	})

	return nil
}

func (s *S3) Stop(ctx context.Context) error {
	return nil
}

func NewS3(cfg *config.Config) *S3 {
	return &S3{
		cfg:    cfg,
		bucket: cfg.BucketName,
	}
}
