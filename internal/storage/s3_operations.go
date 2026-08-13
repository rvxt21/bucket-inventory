package storage

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *S3) UploadObject(ctx context.Context, key string, body io.Reader, contentType string) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	})

	return err
}

func (s *S3) GetObject(ctx context.Context, key string) (io.ReadCloser, *string, error) {
	obj, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    &key,
	})
	if err != nil {
		return nil, nil, err
	}

	return obj.Body, obj.ContentType, nil
}

func (s *S3) PresignGetObject(ctx context.Context, key string, ttl time.Duration) (string, error) {
	req, err := s.presign.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(ttl))
	if err != nil {
		return "", err
	}

	if s.cfg.PublicEndpoint == "" {
		return req.URL, nil
	}

	return rewriteHost(req.URL, s.cfg.PublicEndpoint)
}

func rewriteHost(rawURL, endpoint string) (string, error) {
	signed, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	public, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	signed.Scheme = public.Scheme
	signed.Host = public.Host

	return signed.String(), nil
}
